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
 * int.go: Implementation of Python's integer type
 */

package runtime

import (
	"fmt"
	"strconv"
	"strings"
)

// Int ..
type Int interface {
	Object
	Value() int
}

type intStruct struct {
	Object
	value int
}

func (is *intStruct) Value() int {
	return is.value
}

var _ Int = (*intStruct)(nil)

func newInt(t Type, value int) Int {
	return &intStruct{newObject(t), value}
}

// NewInt ..
func NewInt(value int) Int {
	return newInt(IntType, value)
}

// IntType ..
var IntType = newBuiltinType("int")

func intNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("int", args, kwArgs, 0, ObjectType, ObjectType)
	if t != IntType {
		i := intNew(IntType, args, kwArgs)
		return newInt(t, i.(Int).Value())
	}
	argCount := len(args)
	if argCount == 0 {
		return NewInt(0)
	}
	o := args[0]
	if argCount == 2 {
		if !o.IsInstance(StrType) {
			panic(RaiseType(TypeErrorType, "int() can't convert non-string with explicit base"))
		}
		s := o.(Str).Value()
		base := args[1].(Int).Value()
		if !(base == 0 || (base >= 2 && base <= 36)) {
			panic(RaiseType(ValueErrorType, "int() base must be >= 2 and <= 36, or 0"))
		}
		i, err := intParse(s, base)
		if err != nil {
			panic(RaiseType(ValueErrorType, fmt.Sprintf(
				"invalid literal for int() with base %d: %s",
				base, Repr(o).Value())))
		}
		return NewInt(i)
	}
	// try __int__()
	if intFn := o.Type().Slots().Int; intFn != nil {
		return intFn(o)
	}
	// give up
	panic(RaiseType(TypeErrorType, fmt.Sprintf(
		"int() argument must be a string, a bytes-like object or a number, not '%s'",
		o.Type().Name())))
}

func intParse(s string, base int) (int, error) {
	s = strings.TrimSpace(s)
	if len(s) > 2 && s[0] == '0' {
		switch s[1] {
		case 'b', 'B':
			if base == 0 || base == 2 {
				base = 2
				s = s[2:]
			}
		case 'o', 'O':
			if base == 0 || base == 8 {
				base = 8
				s = s[2:]
			}
		case 'x', 'X':
			if base == 0 || base == 16 {
				base = 16
				s = s[2:]
			}
		default:
			base = 8
		}
	}
	if base == 0 {
		base = 10
	}
	i, err := strconv.ParseInt(s, base, strconv.IntSize)
	return int(i), err
}

func intHash(o Object) Object {
	return o
}

func intRepr(o Object) Object {
	i := o.(Int)
	return NewStr(strconv.Itoa(i.Value()))
}

func intInt(o Object) Object {
	i := o.(Int)
	return NewInt(i.Value())
}

func intFloat(o Object) Object {
	i := o.(Int)
	return NewFloat(float64(i.Value()))
}

func intBool(o Object) Object {
	i := o.(Int)
	return NewBool(i.Value() != 0)
}

func intCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(IntType) {
		return NotImplemented
	}
	i := a.(Int).Value()
	j := b.(Int).Value()
	var r bool
	switch op {
	case compareEq:
		r = (i == j)
	case compareNe:
		r = (i != j)
	case compareLt:
		r = (i < j)
	case compareGt:
		r = (i > j)
	case compareLe:
		r = (i <= j)
	case compareGe:
		r = (i >= j)
	}
	return NewBool(r)
}

func intNeg(o Object) Object {
	i := o.(Int)
	return NewInt(-i.Value())
}

func intPos(o Object) Object {
	return o
}

func intBinaryOp(opt binaryOpType, a, b Object, refl bool) Object {
	if !b.IsInstance(IntType) {
		return NotImplemented
	}
	ai := a.(Int).Value()
	bi := b.(Int).Value()
	if refl {
		ai, bi = bi, ai
	}
	var ci int
	switch opt {
	case binaryOpAdd:
		ci = ai + bi
	case binaryOpSub:
		ci = ai - bi
	case binaryOpMul:
		ci = ai * bi
	case binaryOpMod:
		if bi == 0 {
			panic(RaiseType(ZeroDivisionErrorType,
				"integer division or modulo by zero"))
		}
		ci = ai % bi
	case binaryOpTrueDiv:
		if bi == 0 {
			panic(RaiseType(ZeroDivisionErrorType,
				"division by zero"))
		}
		return NewFloat(float64(ai) / float64(bi))
	case binaryOpFloorDiv:
		if bi == 0 {
			panic(RaiseType(ZeroDivisionErrorType,
				"integer division or modulo by zero"))
		}
		ci = ai / bi
	case binaryOpPow:
		if bi < 0 {
			return NewFloat(1.0 / float64(powInt(ai, -bi)))
		}
		ci = powInt(ai, bi)
	default:
		return NotImplemented
	}
	return NewInt(ci)
}

func powInt(a, n int) int {
	r := 1
	for n > 0 {
		if n%2 == 1 {
			r = r * a
		}
		a = a * a
		n = n / 2
	}
	return r
}

func initIntType(slots *typeSlots, dict map[string]Object) {
	slots.New = intNew
	slots.Hash = intHash
	slots.Repr = intRepr
	slots.Int = intInt
	slots.Float = intFloat
	slots.Bool = intBool
	slots.Eq = makeCompareFn(intCompare, compareEq)
	slots.Ne = makeCompareFn(intCompare, compareNe)
	slots.Lt = makeCompareFn(intCompare, compareLt)
	slots.Gt = makeCompareFn(intCompare, compareGt)
	slots.Le = makeCompareFn(intCompare, compareLe)
	slots.Ge = makeCompareFn(intCompare, compareGe)
	slots.Neg = intNeg
	slots.Pos = intPos
	slots.Add = makeBinaryOpFn(intBinaryOp, binaryOpAdd, false)
	slots.Sub = makeBinaryOpFn(intBinaryOp, binaryOpSub, false)
	slots.Mul = makeBinaryOpFn(intBinaryOp, binaryOpMul, false)
	slots.Mod = makeBinaryOpFn(intBinaryOp, binaryOpMod, false)
	slots.TrueDiv = makeBinaryOpFn(intBinaryOp, binaryOpTrueDiv, false)
	slots.FloorDiv = makeBinaryOpFn(intBinaryOp, binaryOpFloorDiv, false)
	slots.Pow = makeBinaryOpFn(intBinaryOp, binaryOpPow, false)
	slots.RAdd = makeBinaryOpFn(intBinaryOp, binaryOpAdd, true)
	slots.RSub = makeBinaryOpFn(intBinaryOp, binaryOpSub, true)
	slots.RMul = makeBinaryOpFn(intBinaryOp, binaryOpMul, true)
	slots.RMod = makeBinaryOpFn(intBinaryOp, binaryOpMod, true)
	slots.RTrueDiv = makeBinaryOpFn(intBinaryOp, binaryOpTrueDiv, true)
	slots.RFloorDiv = makeBinaryOpFn(intBinaryOp, binaryOpFloorDiv, true)
	slots.RPow = makeBinaryOpFn(intBinaryOp, binaryOpPow, true)
}
