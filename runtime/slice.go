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
 * slice.go: Implementation of Python's slice type
 */

package runtime

// Slice ..
type Slice interface {
	Object
	Start() Object
	Stop() Object
	Step() Object
	CalcSlice(elems int) (start, stop, step, len int)
}

type sliceStruct struct {
	Object
	start Object
	stop  Object
	step  Object
}

func (ss *sliceStruct) Start() Object {
	return ss.start
}

func (ss *sliceStruct) Stop() Object {
	return ss.stop
}

func (ss *sliceStruct) Step() Object {
	return ss.step
}

func (ss *sliceStruct) CalcSlice(elems int) (int, int, int, int) {
	asIndex := func(o Object) int {
		if !o.IsInstance(IntType) {
			panic(RaiseType(TypeErrorType,
				"slice indices must be integers or None or have an __index__ method"))
		}
		return o.(Int).Value()
	}
	// calc step
	var step = 1
	if ss.step != None {
		step = asIndex(ss.step)
		if step == 0 {
			panic(RaiseType(ValueErrorType, "slice step cannot be zero"))
		}
	}
	// defaults for start, stop based on the direction of step
	var startDef, stopDef int
	if step > 0 {
		startDef = 0
		stopDef = elems
	} else {
		startDef = elems - 1
		stopDef = -1
	}
	// calc start
	start := startDef
	if ss.start != None {
		start = seqClampIndex(asIndex(ss.start), elems)
	}
	// calc stop
	stop := stopDef
	if ss.stop != None {
		stop = seqClampIndex(asIndex(ss.stop), elems)
	}
	// calc effective stop and slice length
	stop, len := seqRange(start, stop, step)
	return start, stop, step, len
}

var _ Slice = (*sliceStruct)(nil)

// NewSlice ..
func NewSlice(start, stop, step Object) Slice {
	return &sliceStruct{newObject(SliceType), start, stop, step}
}

// SliceType ..
var SliceType = newBuiltinType("slice")

func sliceNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("slice", args, kwArgs, 1, ObjectType, ObjectType, ObjectType)
	start, stop, step := None, None, None
	if argc := len(args); argc == 1 {
		stop = args[0]
	} else {
		start = args[0]
		stop = args[1]
		if argc == 3 {
			step = args[2]
		}
	}
	return NewSlice(start, stop, step)
}

func sliceRepr(o Object) Object {
	return NewStr("slice" + Repr(asTuple(o)).Value())
}

func sliceCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(SliceType) {
		return NotImplemented
	}
	return seqCompare(op,
		asTuple(a).Elems(),
		asTuple(b).Elems())
}

func asTuple(o Object) Tuple {
	s := o.(Slice)
	return NewTuple(s.Start(), s.Stop(), s.Step())
}

func initSliceType(slots *typeSlots, dict map[string]Object) {
	slots.New = sliceNew
	slots.Hash = unhashableHash
	slots.Repr = sliceRepr
	slots.Eq = makeCompareFn(sliceCompare, compareEq)
	slots.Ne = makeCompareFn(sliceCompare, compareNe)
	slots.Lt = makeCompareFn(sliceCompare, compareLt)
	slots.Gt = makeCompareFn(sliceCompare, compareGt)
	slots.Le = makeCompareFn(sliceCompare, compareLe)
	slots.Ge = makeCompareFn(sliceCompare, compareGe)
}
