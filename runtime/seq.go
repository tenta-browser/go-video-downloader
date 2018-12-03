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
 * seq.go: Common structures and functions of Python's sequences
 */

package runtime

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// Seq ..
type Seq interface {
	Object
	Elems() []Object
	setElems([]Object)
}

type seqStruct struct {
	Object
	elems []Object
}

func (ss *seqStruct) Elems() []Object {
	return ss.elems
}

func (ss *seqStruct) setElems(elems []Object) {
	ss.elems = elems
}

var _ Seq = (*seqStruct)(nil)

func seqNew(t Type, args Args, kwArgs KWArgs) []Object {
	checkFunctionArgsMin(t.Name(), args, kwArgs, 0, ObjectType)
	if len(args) == 0 {
		return nil
	}
	var elems []Object
	seqApply(args[0], func(items []Object) {
		elems = make([]Object, len(items))
		copy(elems, items)
	})
	return elems
}

func seqHash(elems []Object) int {
	h := newHash()
	for _, elem := range elems {
		elemHash := Hash(elem).Value()
		var err error
		if arch32 {
			err = binary.Write(h, binary.BigEndian, int32(elemHash))
		} else {
			err = binary.Write(h, binary.BigEndian, int64(elemHash))
		}
		if err != nil {
			panic(err)
		}
	}
	return sumHash(h)
}

func seqRepr(elems []Object) string {
	b := strings.Builder{}
	for i, o := range elems {
		if i > 0 {
			b.WriteString(", ")
		}
		s := Repr(o)
		b.WriteString(s.Value())
	}
	return b.String()
}

func seqCompare(op compareOp, as, bs []Object) Object {
	compareFn := compareOpFn[op]
	al := len(as)
	bl := len(bs)
	for i := 0; i < al && i < bl; i++ {
		if !IsTrue(Eq(as[i], bs[i])) {
			return compareFn(as[i], bs[i])
		}
	}
	return compareFn(NewInt(al), NewInt(bl))
}

// seqRange calculates actual stop value and range length
func seqRange(start, stop, step int) (int, int) {
	if stop == start || (stop > start) != (step > 0) {
		return start, 0
	}
	if step > 0 {
		stop--
	} else {
		stop++
	}
	n := (stop-start)/step + 1
	return start + n*step, n
}

func seqCheckedIndex(t Type, len, index int) int {
	if index < 0 {
		index = len + index
	}
	if index < 0 || index >= len {
		panic(RaiseType(IndexErrorType, fmt.Sprintf("%s index out of range", t.Name())))
	}
	return index
}

func seqClampIndex(i, len int) int {
	if i < 0 {
		i += len
		if i < 0 {
			i = 0
		}
	}
	if i > len {
		i = len
	}
	return i
}

func seqGetItem(t Type, elems []Object, index Object) (Object, []Object) {
	if index.IsInstance(IntType) {
		i := seqCheckedIndex(t, len(elems), index.(Int).Value())
		return elems[i], nil
	} else if index.IsInstance(SliceType) {
		start, stop, step, len := index.(Slice).CalcSlice(len(elems))
		if step == 1 {
			return nil, elems[start:stop]
		}
		sliced := make([]Object, len)
		i := 0
		for j := start; j != stop; j += step {
			sliced[i] = elems[j]
			i++
		}
		return nil, sliced
	} else {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"%s indices must be integers or slices, not %s",
			t.Name(), index.Type().Name())))
	}
}

func seqForEach(o Object, callback func(Object)) {
	it := Iter(o)
	defer Catch(StopIterationType, nil, nil)
	for {
		callback(Next(it))
	}
}

func seqApply(o Object, callback func([]Object)) {
	switch o.Type() {
	case ListType, TupleType:
		callback(o.(Seq).Elems())
	default:
		items := []Object{}
		seqForEach(o, func(item Object) {
			items = append(items, item)
		})
		callback(items)
	}
}

func seqFindFirst(o Object, pred func(Object) bool) bool {
	for it := Iter(o); ; {
		item := NextDefault(it, AfterLast)
		if item == AfterLast {
			return false
		}
		if pred(item) {
			return true
		}
	}
}

func seqFindElem(elems []Object, o Object) int {
	for i, elem := range elems {
		if IsTrue(Eq(elem, o)) {
			return i
		}
	}
	return -1
}

func seqFindElemBetween(elems []Object, o Object, start, stop int) int {
	len := len(elems)
	start = seqClampIndex(start, len)
	stop = seqClampIndex(stop, len)
	index := -1
	if start < len && start < stop {
		index = seqFindElem(elems[start:stop], o)
	}
	if index >= 0 {
		index += start
	}
	return index
}

func seqAdd(a, b Object) []Object {
	aelems := a.(Seq).Elems()
	belems := b.(Seq).Elems()
	return append(aelems, belems...)
}

func seqMul(a Object, c Object) []Object {
	elems := a.(Seq).Elems()
	cnt := c.(Int).Value()
	res := []Object{}
	for i := 0; i < cnt; i++ {
		res = append(res, elems...)
	}
	return res
}

func seqContains(o Object, needle Object) bool {
	return seqFindFirst(o, func(item Object) bool {
		return IsTrue(Eq(needle, item))
	})
}

// ---- Sequence protocol iteration ----

type seqIterator interface {
	Object
	Seq() Object
	Index() int
	setSeq(Object)
	setIndex(int)
}

type seqIteratorStruct struct {
	Object
	seq   Object
	index int
}

func (sis *seqIteratorStruct) Seq() Object {
	return sis.seq
}

func (sis *seqIteratorStruct) Index() int {
	return sis.index
}

func (sis *seqIteratorStruct) setSeq(seq Object) {
	sis.seq = seq
}

func (sis *seqIteratorStruct) setIndex(index int) {
	sis.index = index
}

var _ seqIterator = (*seqIteratorStruct)(nil)

func newSeqIterator(seq Object) seqIterator {
	return &seqIteratorStruct{newObject(seqIteratorType), seq, 0}
}

var seqIteratorType = newBuiltinType("iterator")

func seqIteratorIter(o Object) Object {
	return o
}

func seqIteratorNext(o Object) Object {
	si := o.(seqIterator)
	seq := si.Seq()
	if seq == nil {
		// iteration already over
		panic(Raise(StopIterationType))
	}
	index := si.Index()
	item := func() Object {
		defer Catch(IndexErrorType, func(ex BaseException) {
			// got IndexError, iteration over from now on
			si.setSeq(nil)
			panic(Raise(StopIterationType))
		}, nil)
		return GetItem(seq, NewInt(index))
	}()
	si.setIndex(index + 1)
	return item
}

func initSeqIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.Iter = seqIteratorIter
	slots.Next = seqIteratorNext
}
