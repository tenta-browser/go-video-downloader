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
 * tuple.go: Implementation of Python's tuple type
 */

package runtime

import (
	"fmt"
)

// Tuple ..
type Tuple interface {
	Seq
}

func newTuple(t Type, elems []Object) Tuple {
	return &seqStruct{newObject(t), elems}
}

// NewTuple ..
func NewTuple(elems ...Object) Tuple {
	return newTuple(TupleType, elems)
}

// TupleType ..
var TupleType = newBuiltinType("tuple")

func tupleNew(t Type, args Args, kwArgs KWArgs) Object {
	elems := seqNew(t, args, kwArgs)
	return newTuple(t, elems)
}

func tupleHash(o Object) Object {
	t := o.(Tuple)
	return NewInt(seqHash(t.Elems()))
}

func tupleRepr(o Object) Object {
	t := o.(Tuple)
	elems := t.Elems()
	s := seqRepr(elems)
	if len(elems) == 1 {
		s = fmt.Sprintf("(%s,)", s)
	} else {
		s = fmt.Sprintf("(%s)", s)
	}
	return NewStr(s)
}

func tupleGetItem(o Object, key Object) Object {
	item, items := seqGetItem(o.Type(), o.(Tuple).Elems(), key)
	if item != nil {
		return item
	}
	return NewTuple(items...)
}

func tupleLen(o Object) Object {
	t := o.(Tuple)
	return NewInt(len(t.Elems()))
}

func tupleCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(TupleType) {
		return NotImplemented
	}
	return seqCompare(op, a.(Tuple).Elems(), b.(Tuple).Elems())
}

func tupleAdd(a, b Object) Object {
	if !b.IsInstance(TupleType) {
		return NotImplemented
	}
	return NewTuple(seqAdd(a, b)...)
}

func tupleMul(o Object, c Object) Object {
	if !c.IsInstance(IntType) {
		return NotImplemented
	}
	return NewTuple(seqMul(o, c)...)
}

func tupleContains(a Object, needle Object) Object {
	return NewBool(seqContains(a, needle))
}

func tupleIndex(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("index", args, kwArgs, 2, TupleType, ObjectType,
		IntType, IntType)
	argc := len(args)
	elems := args[0].(Tuple).Elems()
	value := args[1]
	start, stop := 0, len(elems)
	if argc >= 3 {
		start = args[2].(Int).Value()
	}
	if argc >= 4 {
		stop = args[3].(Int).Value()
	}
	index := seqFindElemBetween(elems, value, start, stop)
	if index < 0 {
		panic(RaiseType(ValueErrorType, fmt.Sprintf(
			"tuple.index(x): %s not in tuple", Repr(value).Value())))
	}
	return NewInt(index)
}

func initTupleType(slots *typeSlots, dict map[string]Object) {
	dict["index"] = newBuiltinFunction("index", tupleIndex)
	slots.New = tupleNew
	slots.Hash = tupleHash
	slots.Repr = tupleRepr
	slots.GetItem = tupleGetItem
	slots.Len = tupleLen
	slots.Eq = makeCompareFn(tupleCompare, compareEq)
	slots.Ne = makeCompareFn(tupleCompare, compareNe)
	slots.Lt = makeCompareFn(tupleCompare, compareLt)
	slots.Gt = makeCompareFn(tupleCompare, compareGt)
	slots.Le = makeCompareFn(tupleCompare, compareLe)
	slots.Ge = makeCompareFn(tupleCompare, compareGe)
	slots.Add = tupleAdd
	slots.Mul = tupleMul
	slots.RMul = tupleMul
	slots.Contains = tupleContains
}
