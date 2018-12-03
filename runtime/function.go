/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * function.go: Implementation of Python's function type and
 *              of other related types
 */

package runtime

import (
	"fmt"
	"strconv"
	"strings"
)

// Param ..
type Param struct {
	Name string
	Def  Object
}

// Args ..
type Args []Object

// NewArgs ...
func NewArgs(args ...Object) Args {
	return args
}

// KWArg ..
type KWArg struct {
	Name  string
	Value Object
}

// KWArgs ..
type KWArgs []KWArg

func (kwArgs KWArgs) get(name string, def Object) Object {
	for _, kwarg := range kwArgs {
		if kwarg.Name == name {
			return kwarg.Value
		}
	}
	return def
}

func (kwArgs KWArgs) unpack() KWArgs {
	unpacked := make(KWArgs, 0, len(kwArgs))
	add := func(kwArg KWArg) {
		if unpacked.get(kwArg.Name, nil) != nil {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"got multiple values for keyword argument '%s'", kwArg.Name)))
		}
		unpacked = append(unpacked, kwArg)
	}
	for _, kwArg := range kwArgs {
		if kwArg.Name != "" {
			// simple kwarg
			add(kwArg)
			continue
		}
		// **kwargs construct, value should be a Str->Object dict
		value := kwArg.Value
		if !value.IsInstance(DictType) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"argument after ** must be a mapping, not %s", value.Type().Name())))
		}
		it := value.(Dict).Iterator()
		for e := it(); e != nil; e = it() {
			key, value := e.Key, e.Value
			if !key.IsInstance(StrType) {
				panic(RaiseType(TypeErrorType, "keywords must be strings"))
			}
			add(KWArg{
				Name:  key.(Str).Value(),
				Value: value,
			})
		}
	}
	return unpacked
}

type fun func([]Object) Object

// Function ..
type Function interface {
	Object
	Name() string
	Params() []Param
	NumKWOnly() int
	HasVarArgs() bool
	HasKWArgs() bool
	Fn() fun
}

type functionStruct struct {
	Object
	name       string
	params     []Param
	numKWOnly  int
	hasVarArgs bool
	hasKWArgs  bool
	fn         fun
}

func (fs *functionStruct) Name() string {
	return fs.name
}

func (fs *functionStruct) Params() []Param {
	return fs.params
}

func (fs *functionStruct) NumKWOnly() int {
	return fs.numKWOnly
}

func (fs *functionStruct) HasVarArgs() bool {
	return fs.hasVarArgs
}

func (fs *functionStruct) HasKWArgs() bool {
	return fs.hasKWArgs
}

func (fs *functionStruct) Fn() fun {
	return fs.fn
}

var _ Function = (*functionStruct)(nil)

// NewFunction ..
func NewFunction(name string, params []Param, numKWOnly int,
	hasVarArgs bool, hasKwArgs bool, fn fun) Function {

	return &functionStruct{
		newObject(FunctionType),
		name,
		params,
		numKWOnly,
		hasVarArgs,
		hasKwArgs,
		fn,
	}
}

// NewSimpleFunction ..
func NewSimpleFunction(name string, parnames []string, fn fun) Function {
	params := make([]Param, len(parnames))
	for i, parname := range parnames {
		params[i] = Param{Name: parname}
	}
	return NewFunction(name, params, 0, false, false, fn)
}

func newBuiltinFunction(name string, fn func(Args, KWArgs) Object) Function {
	return NewFunction(name, nil, 0, true, true, func(args []Object) Object {
		// unpack args tuple
		fnArgs := args[0].(Tuple).Elems()
		// unpack kwargs dict
		var fnKWArgs KWArgs
		it := args[1].(Dict).Iterator()
		for e := it(); e != nil; e = it() {
			fnKWArgs = append(fnKWArgs, KWArg{
				Name:  e.Key.(Str).Value(),
				Value: e.Value,
			})
		}
		return fn(fnArgs, fnKWArgs)
	})
}

// FunctionType ..
var FunctionType = newBuiltinType("function")

func functionCall(o Object, args Args, kwArgs KWArgs) Object {
	f := o.(Function)

	// check if there are too much positional arguments given
	argCount := len(args)
	params := f.Params()
	paramCount := len(params)
	if maxPositionalCount := paramCount - f.NumKWOnly(); argCount > maxPositionalCount {
		if !f.HasVarArgs() {
			minPositionalCount := 0
			for _, param := range params {
				if param.Def != nil {
					break
				}
				minPositionalCount++
			}
			var positionalCountStr string
			if minPositionalCount == maxPositionalCount {
				positionalCountStr = strconv.Itoa(minPositionalCount)
			} else {
				positionalCountStr = fmt.Sprintf("from %d to %d",
					minPositionalCount, maxPositionalCount)
			}
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"%s() takes %s positional arguments but %d %s given",
				f.Name(),
				positionalCountStr,
				argCount,
				plural(argCount, "was", "were"))))
		}
		argCount = maxPositionalCount
	}

	validArgs := make([]Object, paramCount)

	// copy available positional arguments
	paramIdx := 0
	for ; paramIdx < argCount; paramIdx++ {
		validArgs[paramIdx] = args[paramIdx]
	}

	// swallow extra positional arguments into a var-args tuple
	if f.HasVarArgs() {
		validArgs = append(validArgs, NewTuple(args[paramIdx:]...))
	}

	var kwArgsDict Dict
	if f.HasKWArgs() {
		kwArgsDict = NewDict()
		validArgs = append(validArgs, kwArgsDict)
	}

	// insert keyword arguments into the correct position in the argument slice
	for _, kwArg := range kwArgs {
		name := kwArg.Name
		paramIdx := 0
		for paramIdx = 0; paramIdx < paramCount; paramIdx++ {
			if params[paramIdx].Name == name {
				if validArgs[paramIdx] != nil {
					panic(RaiseType(TypeErrorType, fmt.Sprintf(
						"%s() got multiple values for argument '%s'",
						f.Name(), name)))
				}
				validArgs[paramIdx] = kwArg.Value
				break
			}
		}
		if paramIdx == paramCount {
			if !f.HasKWArgs() {
				panic(RaiseType(TypeErrorType, fmt.Sprintf(
					"%s() got an unexpected keyword argument '%s'",
					f.Name(), name)))
			}
			kwArgsDict.SetItem(NewStr(name), kwArg.Value)
		}
	}

	var missing []string

	// apply default values for non-filled arguments
	for ; paramIdx < paramCount; paramIdx++ {
		if validArgs[paramIdx] != nil {
			continue
		}
		param := params[paramIdx]
		if param.Def == nil {
			missing = append(missing, param.Name)
		}
		validArgs[paramIdx] = param.Def
	}

	if l := len(missing); l > 0 {
		// there are still missing arguments, raise error
		b := strings.Builder{}
		fmt.Fprintf(&b, "%s() missing %d required positional %s: ",
			f.Name(), l, plural(l, "argument", "arguments"))
		for i := 0; i < l; i++ {
			if i > 0 {
				if l > 2 {
					b.WriteString(", ")
				} else {
					b.WriteByte(' ')
				}
				if i == l-1 {
					b.WriteString("and ")
				}
			}
			b.WriteByte('\'')
			b.WriteString(missing[i])
			b.WriteByte('\'')
		}
		panic(RaiseType(TypeErrorType, b.String()))
	}

	return f.Fn()(validArgs)
}

func functionGet(desc Object, inst Object, t Type) Object {
	// Note: all functions are descriptors that bind them
	// to the instance, effectively converting them to methods.
	if inst == nil || inst == None {
		return desc
	}
	return Cal(MethodType, inst, desc)
}

func functionRepr(o Object) Object {
	f := o.(Function)
	return NewStr(fmt.Sprintf("<%s %s at %p>", f.Type().Name(), f.Name(), f))
}

func initFunctionType(slots *typeSlots, dict map[string]Object) {
	slots.Call = functionCall
	slots.Get = functionGet
	slots.Repr = functionRepr
}

// ---- Static method ----

type staticMethod interface {
	Object
	Fun() Object
}

type staticMethodStruct struct {
	Object
	fun Object
}

func (sms *staticMethodStruct) Fun() Object {
	return sms.fun
}

var _ staticMethod = (*staticMethodStruct)(nil)

func newStaticMethod(t Type, fun Object) staticMethod {
	return &staticMethodStruct{newObject(t), fun}
}

// StaticMethodType ..
var StaticMethodType = newBuiltinType("staticmethod")

func staticMethodNew(t Type, args Args, kwArgs KWArgs) Object {
	return newStaticMethod(t, args[0])
}

func staticMethodGet(desc Object, inst Object, t Type) Object {
	m := desc.(staticMethod)
	return m.Fun()
}

func initStaticMethodType(slots *typeSlots, dict map[string]Object) {
	slots.New = staticMethodNew
	slots.Get = staticMethodGet
}

// ---- Class method ----

type classMethod interface {
	Object
	Fun() Object
}

type classMethodStruct struct {
	Object
	fun Object
}

func (cms *classMethodStruct) Fun() Object {
	return cms.fun
}

var _ classMethod = (*classMethodStruct)(nil)

func newclassMethod(t Type, fun Object) classMethod {
	return &classMethodStruct{newObject(t), fun}
}

// ClassMethodType ..
var ClassMethodType = newBuiltinType("classmethod")

func classMethodNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("__init__", args, kwArgs, ObjectType)
	return newclassMethod(t, args[0])
}

func classMethodGet(desc Object, inst Object, t Type) Object {
	m := desc.(classMethod)
	return Cal(MethodType, t, m.Fun())
}

func initClassMethodType(slots *typeSlots, dict map[string]Object) {
	slots.New = classMethodNew
	slots.Get = classMethodGet
}
