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
 * float.go: Implementation of Python's float type
 */

package runtime

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

// Float ..
type Float interface {
	Object
	Value() float64
}

type floatStruct struct {
	Object
	value float64
}

func (fs *floatStruct) Value() float64 {
	return fs.value
}

var _ Float = (*floatStruct)(nil)

func newFloat(t Type, value float64) Float {
	return &floatStruct{newObject(t), value}
}

// NewFloat ..
func NewFloat(value float64) Float {
	return newFloat(FloatType, value)
}

// FloatType ..
var FloatType = newBuiltinType("float")

func floatNew(t Type, args Args, kwArgs KWArgs) Object {
	if t != FloatType {
		f := floatNew(FloatType, args, kwArgs)
		return newFloat(t, f.(Float).Value())
	}
	checkFunctionArgsMin("float", args, kwArgs, 0, ObjectType)
	argCount := len(args)
	if argCount == 0 {
		return NewFloat(0)
	}
	o := args[0]
	// try __float__()
	if floatFn := o.Type().Slots().Float; floatFn != nil {
		return floatFn(o)
	}
	// give up
	panic(RaiseType(TypeErrorType, fmt.Sprintf(
		"float() argument must be a string or a number, not '%s'",
		o.Type().Name())))
}

func floatHash(o Object) Object {
	f := o.(Float)
	h := newHash()
	if err := binary.Write(h, binary.BigEndian, f.Value()); err != nil {
		panic(err)
	}
	return NewInt(sumHash(h))
}

func floatRepr(o Object) Object {
	f := o.(Float).Value()
	switch {
	case math.IsInf(f, 1):
		return NewStr("inf")
	case math.IsInf(f, -1):
		return NewStr("-inf")
	case math.IsNaN(f):
		return NewStr("nan")
	default:
		return NewStr(strconv.FormatFloat(f, 'g', 16, 64))
	}
}

func floatInt(o Object) Object {
	f := o.(Float).Value()
	switch {
	case math.IsInf(f, 0):
		panic(RaiseType(OverflowErrorType,
			"cannot convert float infinity to integer"))
	case math.IsNaN(f):
		panic(RaiseType(ValueErrorType,
			"cannot convert float NaN to integer"))
	default:
		return NewInt(int(f))
	}
}

func floatFloat(o Object) Object {
	return o
}

func floatBool(o Object) Object {
	f := o.(Float)
	return NewBool(f.Value() != 0)
}

func floatCompare(op compareOp, a, b Object) Object {
	j, ok := floatCoerce(b)
	if !ok {
		return NotImplemented
	}
	i := a.(Float).Value()
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

func floatCoerce(o Object) (float64, bool) {
	if o.IsInstance(FloatType) {
		return o.(Float).Value(), true
	} else if o.IsInstance(IntType) {
		return float64(o.(Int).Value()), true
	} else {
		return 0, false
	}
}

func floatNeg(o Object) Object {
	f := o.(Float)
	return NewFloat(-f.Value())
}

func floatPos(o Object) Object {
	return o
}

func floatBinaryOp(opt binaryOpType, a, b Object, refl bool) Object {
	bf, ok := floatCoerce(b)
	if !ok {
		return NotImplemented
	}
	af := a.(Float).Value()
	if refl {
		af, bf = bf, af
	}
	var cf float64
	switch opt {
	case binaryOpAdd:
		cf = af + bf
	case binaryOpSub:
		cf = af - bf
	case binaryOpMul:
		cf = af * bf
	case binaryOpMod:
		if bf == 0 {
			panic(RaiseType(ZeroDivisionErrorType, "float modulo"))
		}
		cf = math.Mod(af, bf)
	case binaryOpTrueDiv:
		if bf == 0 {
			panic(RaiseType(ZeroDivisionErrorType,
				"float division by zero"))
		}
		cf = af / bf
	case binaryOpFloorDiv:
		if bf == 0 {
			panic(RaiseType(ZeroDivisionErrorType,
				"float divmod()"))
		}
		cf = math.Floor(af / bf)
	case binaryOpPow:
		cf = math.Pow(af, bf)
	default:
		return NotImplemented
	}
	return NewFloat(cf)
}

func initFloatType(slots *typeSlots, dict map[string]Object) {
	slots.New = floatNew
	slots.Hash = floatHash
	slots.Repr = floatRepr
	slots.Int = floatInt
	slots.Float = floatFloat
	slots.Bool = floatBool
	slots.Eq = makeCompareFn(floatCompare, compareEq)
	slots.Ne = makeCompareFn(floatCompare, compareNe)
	slots.Lt = makeCompareFn(floatCompare, compareLt)
	slots.Gt = makeCompareFn(floatCompare, compareGt)
	slots.Le = makeCompareFn(floatCompare, compareLe)
	slots.Ge = makeCompareFn(floatCompare, compareGe)
	slots.Neg = floatNeg
	slots.Pos = floatPos
	slots.Add = makeBinaryOpFn(floatBinaryOp, binaryOpAdd, false)
	slots.Sub = makeBinaryOpFn(floatBinaryOp, binaryOpSub, false)
	slots.Mul = makeBinaryOpFn(floatBinaryOp, binaryOpMul, false)
	slots.Mod = makeBinaryOpFn(floatBinaryOp, binaryOpMod, false)
	slots.TrueDiv = makeBinaryOpFn(floatBinaryOp, binaryOpTrueDiv, false)
	slots.FloorDiv = makeBinaryOpFn(floatBinaryOp, binaryOpFloorDiv, false)
	slots.Pow = makeBinaryOpFn(floatBinaryOp, binaryOpPow, false)
	slots.RAdd = makeBinaryOpFn(floatBinaryOp, binaryOpAdd, true)
	slots.RSub = makeBinaryOpFn(floatBinaryOp, binaryOpSub, true)
	slots.RMul = makeBinaryOpFn(floatBinaryOp, binaryOpMul, true)
	slots.RMod = makeBinaryOpFn(floatBinaryOp, binaryOpMod, true)
	slots.RTrueDiv = makeBinaryOpFn(floatBinaryOp, binaryOpTrueDiv, true)
	slots.RFloorDiv = makeBinaryOpFn(floatBinaryOp, binaryOpFloorDiv, true)
	slots.RPow = makeBinaryOpFn(floatBinaryOp, binaryOpPow, true)
}
