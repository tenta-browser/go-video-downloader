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
 * base.go: Basic runtime functions used by generated code
 */

package runtime

import (
	"fmt"
	"strings"
)

var (
	// PrintingEnabled ..
	PrintingEnabled = true
)

var (
	// BlockExitNormally ..
	BlockExitNormally = newObject(ObjectType)
	// BlockExitReturn ..
	BlockExitReturn = newObject(ObjectType)
	// BlockExitContinue ..
	BlockExitContinue = newObject(ObjectType)
	// BlockExitBreak ..
	BlockExitBreak = newObject(ObjectType)
)

// Call ..
func Call(o Object, args Args, kwArgs KWArgs) Object {
	call := o.Type().Slots().Call
	if call == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object is not callable", o.Type().Name())))
	}
	return call(o, args, kwArgs.unpack())
}

// Cal ..
func Cal(o Object, args ...Object) Object {
	return Call(o, Args(args), nil)
}

// Raise ..
func Raise(o Object) Object {
	const errMsg = "exceptions must derive from BaseException"
	var ex Object
	if o.IsInstance(TypeType) {
		if !o.(Type).IsSubType(BaseExceptionType) {
			panic(RaiseType(TypeErrorType, errMsg))
		}
		ex = Cal(o)
	} else {
		if !o.IsInstance(BaseExceptionType) {
			panic(RaiseType(TypeErrorType, errMsg))
		}
		ex = o
	}
	return ex
}

// RaiseType ..
func RaiseType(t Type, msg string) Object {
	return Raise(Cal(t, NewStr(msg)))
}

// Catcher ..
type Catcher struct {
	O Object
	H func(BaseException)
}

func catches(o Object, h func(BaseException), ex BaseException) bool {
	var types []Object
	if o.IsInstance(TupleType) {
		types = o.(Tuple).Elems()
	} else {
		types = []Object{o}
	}
	for _, t := range types {
		if !t.IsInstance(TypeType) || !t.(Type).IsSubType(BaseExceptionType) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"catching classes that do not inherit from BaseException is not allowed (%s)",
				t.Type().Name())))
		}
		if ex.IsInstance(t.(Type)) {
			if h != nil {
				h(ex)
			}
			return true
		}
	}
	return false
}

// CatchMulti ..
func CatchMulti(e func(), catchers ...*Catcher) {
	if r := recover(); r != nil {
		if ex, ok := r.(BaseException); ok {
			for _, catcher := range catchers {
				if catches(catcher.O, catcher.H, ex) {
					return
				}
			}
		}
		panic(r)
	}
	if e != nil {
		e()
	}
}

// Catch ..
func Catch(o Object, h func(BaseException), e func()) {
	if r := recover(); r != nil {
		if ex, ok := r.(BaseException); ok {
			if catches(o, h, ex) {
				return
			}
		}
		panic(r)
	}
	if e != nil {
		e()
	}
}

const errGetAttribute = "'%s' object has no attribute '%s'"

// GetAttr ..
func GetAttr(o Object, name string, def Object) (value Object) {
	getAttributeFn := o.Type().Slots().GetAttribute
	if getAttributeFn == nil {
		panic(RaiseType(AttributeErrorType,
			fmt.Sprintf(errGetAttribute, o.Type().Name(), name)))
	}
	defer Catch(AttributeErrorType, func(ex BaseException) {
		if def != nil {
			value = def
		} else {
			panic(Raise(ex))
		}
	}, nil)
	return getAttributeFn(o, NewStr(name))
}

// SetAttr ..
func SetAttr(o Object, name string, value Object) {
	setAttrFn := o.Type().Slots().SetAttr
	if setAttrFn == nil {
		panic(RaiseType(TypeErrorType,
			fmt.Sprintf("can't set attributes of type '%s'", o.Type().Name())))
	}
	setAttrFn(o, NewStr(name), value)
}

// DelAttr ..
func DelAttr(o Object, name string) {
	delAttrFn := o.Type().Slots().DelAttr
	if delAttrFn == nil {
		panic(RaiseType(TypeErrorType,
			fmt.Sprintf("can't delete attributes of type '%s'", o.Type().Name())))
	}
	delAttrFn(o, NewStr(name))
}

// HasAttr ..
func HasAttr(o Object, name string) bool {
	notFound := newObject(ObjectType)
	return GetAttr(o, name, notFound) != notFound
}

// Print ..
func Print(objects []Object, sep Str, end Str) {
	if !PrintingEnabled {
		return
	}
	b := strings.Builder{}
	for idx, elem := range objects {
		if idx > 0 {
			b.WriteString(sep.Value())
		}
		b.WriteString(ToStr(elem).Value())
	}
	b.WriteString(end.Value())
	fmt.Print(b.String())
}

func unhashableHash(o Object) Object {
	panic(RaiseType(TypeErrorType, fmt.Sprintf(
		"unhashable type: '%s'", o.Type().Name())))
}

// Hash ..
func Hash(o Object) Int {
	hashFn := o.Type().Slots().Hash
	if hashFn == nil {
		hashFn = unhashableHash
	}
	return hashFn(o).(Int)
}

// Repr ..
func Repr(o Object) Str {
	reprFn := o.Type().Slots().Repr
	if reprFn == nil {
		tn := o.Type().Name()
		return NewStr(fmt.Sprintf("<%s object at %p>", tn, o))
	}
	return reprFn(o).(Str)
}

// ToStr ..
func ToStr(o Object) Str {
	return Cal(StrType, o).(Str)
}

// ToInt ..
func ToInt(o Object) Int {
	return Cal(IntType, o).(Int)
}

// GetItem ..
func GetItem(o Object, key Object) Object {
	getItemFn := o.Type().Slots().GetItem
	if getItemFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object is not subscriptable", o.Type().Name())))
	}
	return getItemFn(o, key)
}

// SetItem ..
func SetItem(o Object, key Object, value Object) {
	setItemFn := o.Type().Slots().SetItem
	if setItemFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object does not support item assignment", o.Type().Name())))
	}
	setItemFn(o, key, value)
}

// DelItem ..
func DelItem(o Object, key Object) {
	delItemFn := o.Type().Slots().DelItem
	if delItemFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object does not support item deletion", o.Type().Name())))
	}
	delItemFn(o, key)
}

// Len ..
func Len(o Object) Int {
	lenFn := o.Type().Slots().Len
	if lenFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"object of type '%s' has no len()", o.Type().Name())))
	}
	return lenFn(o).(Int)
}

// Iter ..
func Iter(o Object) Object {
	iterFn := o.Type().Slots().Iter
	if iterFn != nil {
		return iterFn(o)
	}
	if o.Type().Slots().GetItem != nil {
		return newSeqIterator(o)
	}
	panic(RaiseType(TypeErrorType, fmt.Sprintf(
		"'%s' object is not iterable", o.Type().Name())))
}

// Next ..
func Next(o Object) Object {
	nextFn := o.Type().Slots().Next
	if nextFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object is not an iterator", o.Type().Name())))
	}
	return nextFn(o)
}

// NextDefault ..
func NextDefault(o Object, def Object) (ret Object) {
	defer Catch(StopIterationType, func(ex BaseException) {
		ret = def
	}, nil)
	return Next(o)
}

// AfterLast is a singleton object used to signal iteration end
// in Next(iter, AfterLast) calls.
var AfterLast = newObject(ObjectType)

// IsTrue ..
func IsTrue(o Object) bool {
	// try with __bool__
	if boolFn := o.Type().Slots().Bool; boolFn != nil {
		return boolFn(o).(Int).Value() != 0
	}
	// try with __len__ != 0
	if lenFn := o.Type().Slots().Len; lenFn != nil {
		len := lenFn(o)
		return len.(Int).Value() != 0
	}
	return true
}

type compareOp int

const (
	compareEq compareOp = iota
	compareNe
	compareLt
	compareGt
	compareLe
	compareGe
)

var compareOpReflection = map[compareOp]compareOp{
	compareEq: compareEq,
	compareNe: compareNe,
	compareLt: compareGt,
	compareGt: compareLt,
	compareLe: compareGe,
	compareGe: compareLe,
}

var compareOpStr = map[compareOp]string{
	compareEq: "==",
	compareNe: "!=",
	compareLt: "<",
	compareGt: ">",
	compareLe: "<=",
	compareGe: ">=",
}

func compare(op compareOp, a, b Object) Object {
	r := tryCompare(op, a, b)
	if r == NotImplemented {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' not supported between instances of '%s' and '%s'",
			compareOpStr[op], a.Type().Name(), b.Type().Name())))
	}
	return r
}

func tryCompare(op compareOp, a, b Object) Object {
	var r Object
	if a.Type() != b.Type() && b.Type().IsSubType(a.Type()) {
		r = tryCompareDirectly(compareOpReflection[op], b, a)
		if r == NotImplemented {
			r = tryCompareDirectly(op, a, b)
		}
	} else {
		r = tryCompareDirectly(op, a, b)
		if r == NotImplemented {
			r = tryCompareDirectly(compareOpReflection[op], b, a)
		}
	}
	if r == NotImplemented {
		// fall back to object identity checking for Eq and Ne
		switch op {
		case compareEq:
			r = NewBool(a == b)
		case compareNe:
			r = NewBool(a != b)
		}
	}
	return r
}

func tryCompareDirectly(op compareOp, a, b Object) Object {
	var compareFn binaryOpFn
	switch slots := a.Type().Slots(); op {
	case compareEq:
		compareFn = slots.Eq
	case compareNe:
		compareFn = slots.Ne
	case compareLt:
		compareFn = slots.Lt
	case compareGt:
		compareFn = slots.Gt
	case compareLe:
		compareFn = slots.Le
	case compareGe:
		compareFn = slots.Ge
	}
	if compareFn == nil {
		return NotImplemented
	}
	return compareFn(a, b)
}

var (
	// Eq ..
	Eq = makeCompareFn(compare, compareEq)
	// Ne ..
	Ne = makeCompareFn(compare, compareNe)
	// Lt ..
	Lt = makeCompareFn(compare, compareLt)
	// Gt ..
	Gt = makeCompareFn(compare, compareGt)
	// Le ..
	Le = makeCompareFn(compare, compareLe)
	// Ge ..
	Ge = makeCompareFn(compare, compareGe)
)

var compareOpFn = map[compareOp]binaryOpFn{
	compareEq: Eq,
	compareNe: Ne,
	compareLt: Lt,
	compareGt: Gt,
	compareLe: Le,
	compareGe: Ge,
}

// Neg ..
func Neg(o Object) Object {
	negFn := o.Type().Slots().Neg
	if negFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"bad operand type for unary -: '%s'", o.Type().Name())))
	}
	return negFn(o)
}

// Pos ..
func Pos(o Object) Object {
	posFn := o.Type().Slots().Pos
	if posFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"bad operand type for unary +: '%s'", o.Type().Name())))
	}
	return posFn(o)
}

func doBinaryOp(opName string, op, rop binaryOpFn, a, b Object) Object {
	var r Object
	if a.Type() != b.Type() && b.Type().IsSubType(a.Type()) {
		r = tryDoBinaryOp(rop, b, a)
		if r == NotImplemented {
			r = tryDoBinaryOp(op, a, b)
		}
	} else {
		r = tryDoBinaryOp(op, a, b)
		if r == NotImplemented {
			r = tryDoBinaryOp(rop, b, a)
		}
	}
	if r == NotImplemented {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"unsupported operand type(s) for %s: '%s' and '%s'",
			opName, a.Type().Name(), b.Type().Name())))
	}
	return r
}

func tryDoBinaryOp(op binaryOpFn, a, b Object) Object {
	if op == nil {
		return NotImplemented
	}
	return op(a, b)
}

type binaryOpType int

const (
	binaryOpAdd binaryOpType = iota
	binaryOpSub
	binaryOpMul
	binaryOpMod
	binaryOpTrueDiv
	binaryOpFloorDiv
	binaryOpPow
	binaryOpOr
	binaryOpXor
	binaryOpAnd
	binaryOpLShift
	binaryOpRShift
)

// Add ..
func Add(a, b Object) Object {
	return doBinaryOp("+", a.Type().Slots().Add, b.Type().Slots().RAdd, a, b)
}

// Sub ..
func Sub(a, b Object) Object {
	return doBinaryOp("-", a.Type().Slots().Sub, b.Type().Slots().RSub, a, b)
}

// Mul ..
func Mul(a, b Object) Object {
	return doBinaryOp("*", a.Type().Slots().Mul, b.Type().Slots().RMul, a, b)
}

// Mod ..
func Mod(a, b Object) Object {
	return doBinaryOp("%", a.Type().Slots().Mod, b.Type().Slots().RMod, a, b)
}

// TrueDiv ..
func TrueDiv(a, b Object) Object {
	return doBinaryOp("/", a.Type().Slots().TrueDiv, b.Type().Slots().RTrueDiv, a, b)
}

// FloorDiv ..
func FloorDiv(a, b Object) Object {
	return doBinaryOp("//", a.Type().Slots().FloorDiv, b.Type().Slots().RFloorDiv, a, b)
}

// Pow ..
func Pow(a, b Object) Object {
	return doBinaryOp("**", a.Type().Slots().Pow, b.Type().Slots().RPow, a, b)
}

// Or ..
func Or(a, b Object) Object {
	return doBinaryOp("|", a.Type().Slots().Or, b.Type().Slots().ROr, a, b)
}

// Xor ..
func Xor(a, b Object) Object {
	return doBinaryOp("^", a.Type().Slots().Xor, b.Type().Slots().RXor, a, b)
}

// And ..
func And(a, b Object) Object {
	return doBinaryOp("&", a.Type().Slots().And, b.Type().Slots().RAnd, a, b)
}

// LShift ..
func LShift(a, b Object) Object {
	return doBinaryOp("<<", a.Type().Slots().LShift, b.Type().Slots().RLShift, a, b)
}

// RShift ..
func RShift(a, b Object) Object {
	return doBinaryOp(">>", a.Type().Slots().RShift, b.Type().Slots().RRShift, a, b)
}

func doInplaceOp(iop, op binaryOpFn, a, b Object) Object {
	if iop != nil {
		return iop(a, b)
	}
	return op(a, b)
}

// IAdd ..
func IAdd(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().IAdd, Add, a, b)
}

// ISub ..
func ISub(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().ISub, Sub, a, b)
}

// IMul ..
func IMul(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().IMul, Mul, a, b)
}

// IMod ..
func IMod(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().IMod, Mod, a, b)
}

// ITrueDiv ..
func ITrueDiv(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().ITrueDiv, TrueDiv, a, b)
}

// IFloorDiv ..
func IFloorDiv(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().IFloorDiv, FloorDiv, a, b)
}

// IPow ..
func IPow(a, b Object) Object {
	return doInplaceOp(a.Type().Slots().IPow, Pow, a, b)
}

// Contains ..
func Contains(o Object, needle Object) bool {
	// try __contains__
	containsFn := o.Type().Slots().Contains
	if containsFn != nil {
		return IsTrue(containsFn(o, needle))
	}
	// try to find by iterating
	return seqContains(o, needle)
}

// IsInstance ..
func IsInstance(o Object, classinfo Object) bool {
	return _isSubclass("isinstance", o.Type(), classinfo)
}

// IsSubclass ..
func IsSubclass(o Object, classinfo Object) bool {
	return _isSubclass("issubclass", o, classinfo)
}

func _isSubclass(fname string, o Object, classinfo Object) bool {
	if !o.IsInstance(TypeType) {
		panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() arg 1 must be a class", fname)))
	}
	t := o.(Type)
	if classinfo.IsInstance(TypeType) {
		return t.IsSubType(classinfo.(Type))
	}
	errMsg := fmt.Sprintf("%s() arg 2 must be a class or tuple of classes", fname)
	if !classinfo.IsInstance(TupleType) {
		panic(RaiseType(TypeErrorType, errMsg))
	}
	for _, elem := range classinfo.(Tuple).Elems() {
		if !elem.IsInstance(TypeType) {
			panic(RaiseType(TypeErrorType, errMsg))
		}
		if t.IsSubType(elem.(Type)) {
			return true
		}
	}
	return false
}

// UnwrapNative ..
func UnwrapNative(o Object) interface{} {
	if !o.IsInstance(NativeType) {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'%s' object is not native", o.Type().Name())))
	}
	return o.(Native).Value()
}

// Placeholder ..
func Placeholder(name string, reason string) Object {
	name = fmt.Sprintf("placeholder %s", name)
	if reason != "" {
		name += fmt.Sprintf(" (reason: %s)", reason)
	}
	return Cal(Cal(TypeType, NewStr(name), NewTuple(), NewDict()))
}

func checkFunctionArgsMin(fname string, args Args, kwArgs KWArgs, minArgs int, types ...Type) {
	expArgs := len(types)
	numArgs := len(args)
	if minArgs == expArgs {
		if numArgs != expArgs {
			panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() expected exactly %d arguments, got %d",
				fname, minArgs, numArgs)))
		}
	} else {
		if numArgs < minArgs {
			panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() expected at least %d arguments, got %d",
				fname, minArgs, numArgs)))
		}
		if numArgs > expArgs {
			panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() expected at most %d arguments, got %d",
				fname, expArgs, numArgs)))
		}
	}
	for i, arg := range args {
		if t := types[i]; !arg.IsInstance(t) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() argument %d must be %s, not %s",
				fname, i+1, t.Name(), arg.Type().Name())))
		}
	}
	if len(kwArgs) != 0 {
		panic(RaiseType(TypeErrorType, fmt.Sprintf("%s() takes no keyword arguments", fname)))
	}
}

func checkFunctionArgs(fname string, args Args, kwArgs KWArgs, types ...Type) {
	checkFunctionArgsMin(fname, args, kwArgs, len(types), types...)
}

func checkFunctionVarArgs(fname string, args Args, kwArgs KWArgs, types ...Type) {
	// TODO error messages are incorrect
	if len(args) <= len(types) {
		checkFunctionArgs(fname, args, kwArgs, types...)
	}
	checkFunctionArgs(fname, args[:len(types)], kwArgs, types...)
}

func checkFunctionKWArgs(fname string, kwArgs KWArgs, expectedTypes map[string]Type) {
	for _, kwArg := range kwArgs {
		name := kwArg.Name
		value := kwArg.Value
		expType, found := expectedTypes[name]
		if !found {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"'%s' is an invalid keyword argument for %s()",
				name, fname)))
		}
		if !value.IsInstance(expType) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"'%s' must be a %s, not %s for %s()",
				name, expType.Name(), value.Type().Name(), fname)))
		}
	}
}
