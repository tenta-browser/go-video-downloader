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
 * slots.go: Definition of Python's slots
 */

package runtime

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	slotsType = reflect.TypeOf(typeSlots{})
	slotCount = slotsType.NumField()
	slotNames = calcSlotNames()
)

type (
	methodFn       func(Object, Args, KWArgs) Object
	getAttributeFn func(Object, Str) Object
	setAttrFn      func(Object, Str, Object)
	delAttrFn      func(Object, Str)
	getFn          func(Object, Object, Type) Object
	setFn          func(Object, Object, Object)
	deleteFn       func(Object, Object)
	unaryOpFn      func(Object) Object
	binaryOpFn     func(Object, Object) Object
	setItemFn      func(Object, Object, Object)
	delItemFn      func(Object, Object)
)

type typeSlots struct {
	// type (TODO add slot wrapping/unwrapping for New)
	New  func(Type, Args, KWArgs) Object
	Init methodFn
	Call methodFn

	// attribute access (TODO GetAttr!)
	GetAttribute getAttributeFn
	SetAttr      setAttrFn
	DelAttr      delAttrFn

	// descriptor protocol
	Get    getFn
	Set    setFn
	Delete deleteFn

	// conversion
	Hash  unaryOpFn
	Repr  unaryOpFn
	Str   unaryOpFn
	Int   unaryOpFn
	Float unaryOpFn
	Bool  unaryOpFn

	// sequence
	GetItem binaryOpFn
	SetItem setItemFn
	DelItem delItemFn
	Len     unaryOpFn

	// iterator
	Iter unaryOpFn
	Next unaryOpFn

	// comparison
	Eq binaryOpFn
	Ne binaryOpFn
	Lt binaryOpFn
	Gt binaryOpFn
	Le binaryOpFn
	Ge binaryOpFn

	// unary ops
	Neg unaryOpFn
	Pos unaryOpFn

	// binary ops
	Add       binaryOpFn
	Sub       binaryOpFn
	Mul       binaryOpFn
	Mod       binaryOpFn
	TrueDiv   binaryOpFn
	FloorDiv  binaryOpFn
	Pow       binaryOpFn
	Or        binaryOpFn
	Xor       binaryOpFn
	And       binaryOpFn
	LShift    binaryOpFn
	RShift    binaryOpFn
	RAdd      binaryOpFn
	RSub      binaryOpFn
	RMul      binaryOpFn
	RMod      binaryOpFn
	RTrueDiv  binaryOpFn
	RFloorDiv binaryOpFn
	RPow      binaryOpFn
	ROr       binaryOpFn
	RXor      binaryOpFn
	RAnd      binaryOpFn
	RLShift   binaryOpFn
	RRShift   binaryOpFn
	IAdd      binaryOpFn
	ISub      binaryOpFn
	IMul      binaryOpFn
	IMod      binaryOpFn
	ITrueDiv  binaryOpFn
	IFloorDiv binaryOpFn
	IPow      binaryOpFn
	IOr       binaryOpFn
	IXor      binaryOpFn
	IAnd      binaryOpFn
	ILShift   binaryOpFn
	IRShift   binaryOpFn

	// misc
	Contains binaryOpFn
	Reversed unaryOpFn
}

func calcSlotNames() []string {
	names := make([]string, slotCount, slotCount)
	for i := 0; i < slotCount; i++ {
		field := slotsType.Field(i)
		names[i] = fmt.Sprintf("__%s__", strings.ToLower(field.Name))
	}
	return names
}

// makeFunctions creates Python-callable functions from the populated slots
func (slots *typeSlots) makeFunctions(t Type) map[string]Function {
	d := make(map[string]Function)
	slotsValue := reflect.ValueOf(slots).Elem()
	for i := 0; i < slotCount; i++ {
		slotName := slotNames[i]
		slotField := slotsValue.Field(i)
		if slotField.IsNil() {
			continue
		}
		var funcFn func(Args, KWArgs) Object
		switch slotFn := slotField.Interface().(type) {
		case methodFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				var checkArgs Args
				if len(args) > 0 {
					checkArgs = args[:1]
				} else {
					checkArgs = args
				}
				checkFunctionArgs(slotName, checkArgs, nil, t)
				return slotFn(args[0], args[1:], kwArgs)
			}
		case getAttributeFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, StrType)
				return slotFn(args[0], args[1].(Str))
			}
		case setAttrFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, StrType, ObjectType)
				slotFn(args[0], args[1].(Str), args[2])
				return None
			}
		case getFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, ObjectType, TypeType)
				return slotFn(args[0], args[1], args[2].(Type))
			}
		case setFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, ObjectType, ObjectType)
				slotFn(args[0], args[1], args[2])
				return None
			}
		case deleteFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, ObjectType)
				slotFn(args[0], args[1])
				return None
			}
		case unaryOpFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t)
				return slotFn(args[0])
			}
		case binaryOpFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, ObjectType)
				return slotFn(args[0], args[1])
			}
		case setItemFn:
			funcFn = func(args Args, kwArgs KWArgs) Object {
				checkFunctionArgs(slotName, args, kwArgs, t, ObjectType, ObjectType)
				slotFn(args[0], args[1], args[2])
				return None
			}
		}
		if funcFn != nil {
			d[slotName] = newBuiltinFunction(slotName, funcFn)
		}
	}
	return d
}

func (slots *typeSlots) fillSlots(dict Dict) {
	slotsValue := reflect.ValueOf(slots).Elem()
	for i := 0; i < slotCount; i++ {
		slotName := slotNames[i]
		dictFunc := dict.GetItem(NewStr(slotName))
		if dictFunc == nil {
			continue
		}
		slotField := slotsValue.Field(i)
		slotIsType := func(h interface{}) bool {
			return slotField.Type().AssignableTo(reflect.TypeOf(h))
		}
		var slotFn interface{}
		usedSymbol := func() {
			usedSymbolDictAttr(dict, slotName)
		}
		switch {
		case slotIsType(methodFn(nil)):
			slotFn = methodFn(func(o Object, args Args, kwArgs KWArgs) Object {
				usedSymbol()
				callArgs := Args{o}
				callArgs = append(callArgs, args...)
				return Call(dictFunc, callArgs, kwArgs)
			})
		case slotIsType(getAttributeFn(nil)):
			slotFn = getAttributeFn(func(o Object, name Str) Object {
				usedSymbol()
				return Cal(dictFunc, o, name)
			})
		case slotIsType(setAttrFn(nil)):
			slotFn = setAttrFn(func(o Object, name Str, value Object) {
				usedSymbol()
				Cal(dictFunc, o, name, value)
			})
		case slotIsType(getFn(nil)):
			slotFn = getFn(func(desc Object, inst Object, t Type) Object {
				usedSymbol()
				return Cal(dictFunc, desc, inst, t)
			})
		case slotIsType(setFn(nil)):
			slotFn = setFn(func(desc Object, inst Object, value Object) {
				usedSymbol()
				Cal(dictFunc, desc, inst, value)
			})
		case slotIsType(deleteFn(nil)):
			slotFn = deleteFn(func(desc Object, inst Object) {
				usedSymbol()
				Cal(dictFunc, desc, inst)
			})
		case slotIsType(unaryOpFn(nil)):
			slotFn = unaryOpFn(func(o Object) Object {
				usedSymbol()
				return Cal(dictFunc, o)
			})
		case slotIsType(binaryOpFn(nil)):
			slotFn = binaryOpFn(func(a, b Object) Object {
				usedSymbol()
				return Cal(dictFunc, a, b)
			})
		case slotIsType(setItemFn(nil)):
			slotFn = setItemFn(func(o Object, key Object, value Object) {
				usedSymbol()
				Cal(o, key, value)
			})
		}
		slotField.Set(reflect.ValueOf(slotFn))
	}
}
