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
 * list.go: Implementation of Python's list type
 */

package runtime

import (
	"fmt"
	"sort"
)

// List ..
type List interface {
	Seq
}

func newList(t Type, elems []Object) List {
	return &seqStruct{newObject(t), elems}
}

// NewList ..
func NewList(elems ...Object) List {
	return newList(ListType, elems)
}

// ListType ..
var ListType = newBuiltinType("list")

func listNew(t Type, args Args, kwArgs KWArgs) Object {
	elems := seqNew(t, args, kwArgs)
	return newList(t, elems)
}

func listHash(o Object) Object {
	l := o.(List)
	return NewInt(seqHash(l.Elems()))
}

func listRepr(o Object) Object {
	l := o.(List)
	elems := l.Elems()
	s := seqRepr(elems)
	return NewStr(fmt.Sprintf("[%s]", s))
}

func listGetItem(o Object, key Object) Object {
	item, items := seqGetItem(o.Type(), o.(List).Elems(), key)
	if item != nil {
		return item
	}
	return NewList(items...)
}

func listSetItem(o Object, key Object, value Object) {
	l := o.(List)
	elems := l.Elems()
	listLen := len(elems)
	if key.IsInstance(IntType) {
		index := key.(Int).Value()
		index = seqCheckedIndex(o.Type(), listLen, index)
		elems[index] = value
	} else if key.IsInstance(SliceType) {
		start, stop, step, sliceLen := key.(Slice).CalcSlice(listLen)
		seqApply(value, func(valElems []Object) {
			valLen := len(valElems)
			if step == 1 {
				newElems := make([]Object, 0, listLen-sliceLen+valLen)
				newElems = append(newElems, elems[0:start]...)
				newElems = append(newElems, valElems...)
				newElems = append(newElems, elems[stop:]...)
				l.setElems(newElems)
			} else if sliceLen == valLen {
				i := 0
				for j := start; j != stop; j += step {
					elems[j] = valElems[i]
					i++
				}
			} else {
				panic(RaiseType(ValueErrorType, fmt.Sprintf(
					"attempt to assign sequence of size %d to extended slice of size %d",
					valLen, sliceLen)))
			}
		})
	} else {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"%s indices must be integers or slices, not %s",
			o.Type().Name(), key.Type().Name())))
	}
}

func listDelItem(o Object, key Object) {
	// Transform: del list[slice] -> list[slice] = []
	if key.IsInstance(IntType) {
		// Transform: [index] -> [index:index+1]
		index := key.(Int).Value()
		key = NewSlice(key, NewInt(index+1), None)
	}
	listSetItem(o, key, NewList())
}

func listLen(o Object) Object {
	l := o.(List)
	return NewInt(len(l.Elems()))
}

func listCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(ListType) {
		return NotImplemented
	}
	return seqCompare(op, a.(List).Elems(), b.(List).Elems())
}

func listAdd(a, b Object) Object {
	if !b.IsInstance(ListType) {
		return NotImplemented
	}
	return NewList(seqAdd(a, b)...)
}

func listMul(o Object, c Object) Object {
	if !c.IsInstance(IntType) {
		return NotImplemented
	}
	return NewList(seqMul(o, c)...)
}

func listIAdd(a, b Object) Object {
	alist := a.(List)
	seqApply(b, func(belems []Object) {
		alist.setElems(append(alist.Elems(), belems...))
	})
	return a
}

func listIMul(o Object, c Object) Object {
	if !c.IsInstance(IntType) {
		return NotImplemented
	}
	o.(List).setElems(seqMul(o, c))
	return o
}

func listContains(a Object, needle Object) Object {
	return NewBool(seqContains(a, needle))
}

func listAppend(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("append", args, kwArgs, ListType, ObjectType)
	l := args[0].(List)
	e := args[1]
	l.setElems(append(l.Elems(), e))
	return None
}

func listExtend(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("extend", args, kwArgs, ListType, ObjectType)
	listIAdd(args[0], args[1])
	return None
}

func listIndex(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("index", args, kwArgs, 2, ListType, ObjectType,
		IntType, IntType)
	argc := len(args)
	elems := args[0].(List).Elems()
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
			"%s is not in the list", Repr(value).Value())))
	}
	return NewInt(index)
}

func listSort(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("sort", args, nil, ListType)
	checkFunctionKWArgs("sort", kwArgs, map[string]Type{
		"key":     ObjectType,
		"reverse": IntType,
	})
	elems := args[0].(List).Elems()
	var keyFn unaryOpFn
	if key := kwArgs.get("key", None); key != None {
		keyFn = func(o Object) Object {
			return Cal(key, o)
		}
	}
	reverse := IsTrue(kwArgs.get("reverse", False))
	sort.SliceStable(elems, func(i, j int) bool {
		var a, b Object
		if reverse {
			a, b = elems[j], elems[i]
		} else {
			a, b = elems[i], elems[j]
		}
		if keyFn != nil {
			a, b = keyFn(a), keyFn(b)
		}
		return IsTrue(Lt(a, b))
	})
	return None
}

func initListType(slots *typeSlots, dict map[string]Object) {
	dict["append"] = newBuiltinFunction("append", listAppend)
	dict["extend"] = newBuiltinFunction("extend", listExtend)
	dict["index"] = newBuiltinFunction("index", listIndex)
	dict["sort"] = newBuiltinFunction("sort", listSort)
	slots.New = listNew
	slots.Hash = listHash
	slots.Repr = listRepr
	slots.GetItem = listGetItem
	slots.SetItem = listSetItem
	slots.DelItem = listDelItem
	slots.Len = listLen
	slots.Eq = makeCompareFn(listCompare, compareEq)
	slots.Ne = makeCompareFn(listCompare, compareNe)
	slots.Lt = makeCompareFn(listCompare, compareLt)
	slots.Gt = makeCompareFn(listCompare, compareGt)
	slots.Le = makeCompareFn(listCompare, compareLe)
	slots.Ge = makeCompareFn(listCompare, compareGe)
	slots.Add = listAdd
	slots.Mul = listMul
	slots.RMul = listMul
	slots.IAdd = listIAdd
	slots.IMul = listIMul
	slots.Contains = listContains
}
