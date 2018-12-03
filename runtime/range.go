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
 * range.go: Implementation of Python's range type
 */

package runtime

import "fmt"

// Range ..
type Range interface {
	Object
	Start() int
	Stop() int
	Step() int
	len() int
}

type rangeStruct struct {
	Object
	start int
	stop  int
	step  int
}

func (rs *rangeStruct) Start() int {
	return rs.start
}

func (rs *rangeStruct) Stop() int {
	return rs.stop
}

func (rs *rangeStruct) Step() int {
	return rs.step
}

func (rs *rangeStruct) len() int {
	return (rs.stop - rs.start) / rs.step
}

var _ Range = (*rangeStruct)(nil)

func newRange(start, stop, step int) Range {
	return &rangeStruct{newObject(RangeType), start, stop, step}
}

// RangeType ..
var RangeType = newBuiltinType("range")

func rangeNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("range", args, kwArgs, 1, IntType, IntType, IntType)
	start, stop, step := 0, 0, 1
	if argc := len(args); argc == 1 {
		stop = args[0].(Int).Value()
	} else {
		start = args[0].(Int).Value()
		stop = args[1].(Int).Value()
		if argc == 3 {
			step = args[2].(Int).Value()
		}
	}
	if step == 0 {
		panic(RaiseType(ValueErrorType, "range() arg 3 must not be zero"))
	}
	stop, _ = seqRange(start, stop, step)
	return newRange(start, stop, step)
}

func rangeRepr(o Object) Object {
	r := o.(Range)
	var s string
	if r.Step() == 1 {
		s = fmt.Sprintf("range(%d, %d)", r.Start(), r.Stop())
	} else {
		s = fmt.Sprintf("range(%d, %d, %d)", r.Start(), r.Stop(), r.Step())
	}
	return NewStr(s)
}

func rangeGetItem(o Object, key Object) Object {
	r := o.(Range)
	index := key.(Int).Value()
	index = seqCheckedIndex(RangeType, r.len(), index)
	return NewInt(r.Start() + r.Step()*index)
}

func rangeLen(o Object) Object {
	r := o.(Range)
	return NewInt(r.len())
}

func rangeIter(o Object) Object {
	r := o.(Range)
	return newRangeIterator(r.Start(), r.Stop(), r.Step())
}

func initRangeType(slots *typeSlots, dict map[string]Object) {
	slots.New = rangeNew
	slots.Repr = rangeRepr
	slots.GetItem = rangeGetItem
	slots.Len = rangeLen
	slots.Iter = rangeIter
}

// ---- RangeIterator ----

// RangeIterator ..
type RangeIterator interface {
	Object
	Curr() int
	Stop() int
	Step() int
	next()
}

type rangeIteratorStruct struct {
	Object
	curr int
	stop int
	step int
}

func (ris *rangeIteratorStruct) Curr() int {
	return ris.curr
}

func (ris *rangeIteratorStruct) Stop() int {
	return ris.stop
}

func (ris *rangeIteratorStruct) Step() int {
	return ris.step
}

func (ris *rangeIteratorStruct) next() {
	ris.curr += ris.step
}

var _ RangeIterator = (*rangeIteratorStruct)(nil)

func newRangeIterator(start, stop, step int) RangeIterator {
	return &rangeIteratorStruct{newObject(RangeIteratorType), start, stop, step}
}

// RangeIteratorType ..
var RangeIteratorType = newBuiltinType("range_iterator")

func rangeIteratorIter(o Object) Object {
	return o
}

func rangeIteratorNext(o Object) Object {
	ri := o.(RangeIterator)
	curr := ri.Curr()
	if ri.Curr() == ri.Stop() {
		panic(Raise(StopIterationType))
	}
	ret := NewInt(curr)
	ri.next()
	return ret
}

func initRangeIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.Iter = rangeIteratorIter
	slots.Next = rangeIteratorNext
}
