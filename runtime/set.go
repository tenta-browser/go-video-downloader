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
 * set.go: Implementation of Python's set type
 */

package runtime

import "fmt"

// Set ..
type Set interface {
	Object
	SDict() Dict

	Add(Object)
	Update(Object)
	Contains(Object) bool
}

type setStruct struct {
	Object
	sdict Dict
}

func (ss *setStruct) SDict() Dict {
	return ss.sdict
}

func (ss *setStruct) Add(o Object) {
	ss.sdict.SetItem(o, None)
}

func (ss *setStruct) Update(o Object) {
	seqForEach(o, func(item Object) {
		ss.Add(item)
	})
}

func (ss *setStruct) Contains(o Object) bool {
	return ss.sdict.GetItem(o) != nil
}

var _ Set = (*setStruct)(nil)

func newSet(t Type) Set {
	return &setStruct{newObject(t), NewDict()}
}

// NewSet ..
func NewSet(elems ...Object) Set {
	s := newSet(SetType)
	for _, elem := range elems {
		s.SDict().SetItem(elem, None)
	}
	return s
}

// SetType ..
var SetType = newBuiltinType("set")

func setNew(t Type, args Args, kwArgs KWArgs) Object {
	return newSet(t)
}

func setInit(o Object, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("set", args, kwArgs, 0, ObjectType)
	s := o.(Set)
	if len(args) > 0 {
		s.Update(args[0])
	}
	return None
}

func setRepr(o Object) Object {
	s := o.(Set)
	elems := s.SDict().Keys()
	len := len(elems)
	str := seqRepr(elems)
	if len > 0 {
		str = fmt.Sprintf("{%s}", str)
	}
	if len == 0 || s.Type() != SetType {
		str = fmt.Sprintf("%s(%s)", s.Type().Name(), str)
	}
	return NewStr(str)
}

func setLen(o Object) Object {
	s := o.(Set)
	return NewInt(s.SDict().Len())
}

func setIter(o Object) Object {
	s := o.(Set)
	return Iter(s.SDict())
}

func setCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(SetType) {
		return NotImplemented
	}
	if op == compareGt || op == compareGe {
		op = compareOpReflection[op]
		a, b = b, a
	}
	as := a.(Set)
	bs := b.(Set)
	alen := as.SDict().Len()
	blen := bs.SDict().Len()
	result := op != compareNe
	switch op {
	case compareLt:
		if alen >= blen {
			return False
		}
	case compareLe:
		if alen > blen {
			return False
		}
	default: // compareEq, compareNe
		if alen != blen {
			return NewBool(!result)
		}
	}
	aiter := as.SDict().Iterator()
	for item := aiter(); item != nil; item = aiter() {
		if !bs.Contains(item.Key) {
			result = !result
			break
		}
	}
	return NewBool(result)
}

func setContains(o Object, needle Object) Object {
	s := o.(Set)
	return NewBool(s.Contains(needle))
}

func setAdd(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("add", args, kwArgs, SetType, ObjectType)
	s := args[0].(Set)
	elem := args[1]
	s.Add(elem)
	return None
}

func initSetType(slots *typeSlots, dict map[string]Object) {
	dict["add"] = newBuiltinFunction("add", setAdd)
	slots.New = setNew
	slots.Init = setInit
	slots.Hash = unhashableHash
	slots.Repr = setRepr
	slots.Len = setLen
	slots.Iter = setIter
	slots.Eq = makeCompareFn(setCompare, compareEq)
	slots.Ne = makeCompareFn(setCompare, compareNe)
	slots.Lt = makeCompareFn(setCompare, compareLt)
	slots.Gt = makeCompareFn(setCompare, compareGt)
	slots.Le = makeCompareFn(setCompare, compareLe)
	slots.Ge = makeCompareFn(setCompare, compareGe)
	slots.Contains = setContains
}
