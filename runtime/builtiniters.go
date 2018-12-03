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
 * builtiniters.go: Various iterators behind Python's builtin functions
 */

package runtime

import "fmt"

// ---- zip ----

type zipIterator interface {
	Object
	Iterators() []Object
	setIterators([]Object)
}

type zipIteratorStruct struct {
	Object
	iterators []Object
}

func (zis *zipIteratorStruct) Iterators() []Object {
	return zis.iterators
}

func (zis *zipIteratorStruct) setIterators(iterators []Object) {
	zis.iterators = iterators
}

var _ zipIterator = (*zipIteratorStruct)(nil)

func newZipIterator(t Type, iterators ...Object) zipIterator {
	return &zipIteratorStruct{newObject(t), iterators}
}

// ZipIteratorType ..
var ZipIteratorType = newBuiltinType("zip")

func zipIteratorNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionVarArgs("zip", args, kwArgs)
	iterators := make([]Object, len(args))
	for i, arg := range args {
		iterators[i] = Iter(arg)
	}
	return newZipIterator(t, iterators...)
}

func zipIteratorIter(o Object) Object {
	return o
}

func zipIteratorNext(o Object) Object {
	zi := o.(zipIterator)
	iterators := zi.Iterators()
	if len(iterators) == 0 {
		panic(Raise(StopIterationType))
	}
	elems := make([]Object, len(iterators))
	for i, iterator := range iterators {
		elem := NextDefault(iterator, AfterLast)
		if elem == AfterLast {
			zi.setIterators(nil)
			panic(Raise(StopIterationType))
		}
		elems[i] = elem
	}
	return NewTuple(elems...)
}

func initZipIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.New = zipIteratorNew
	slots.Iter = zipIteratorIter
	slots.Next = zipIteratorNext
}

// ---- map ----

type mapIterator interface {
	Object
	Callable() Object
	ArgsIter() Object
}

type mapIteratorStruct struct {
	Object
	callable Object
	argsIter Object
}

func (mis *mapIteratorStruct) Callable() Object {
	return mis.callable
}

func (mis *mapIteratorStruct) ArgsIter() Object {
	return mis.argsIter
}

var _ mapIterator = (*mapIteratorStruct)(nil)

func newMapIterator(t Type, callable, argsIter Object) mapIterator {
	return &mapIteratorStruct{newObject(t), callable, argsIter}
}

// MapIteratorType ..
var MapIteratorType = newBuiltinType("map")

func mapIteratorNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionVarArgs("map", args, kwArgs, ObjectType, ObjectType)
	callable := args[0]
	argsIter := Cal(ZipIteratorType, args[1:]...)
	return newMapIterator(t, callable, argsIter)
}

func mapIteratorIter(o Object) Object {
	return o
}

func mapIteratorNext(o Object) Object {
	mi := o.(mapIterator)
	args := NextDefault(mi.ArgsIter(), AfterLast)
	if args == AfterLast {
		panic(Raise(StopIterationType))
	}
	callable := mi.Callable()
	return Cal(callable, args.(Tuple).Elems()...)
}

func initMapIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.New = mapIteratorNew
	slots.Iter = mapIteratorIter
	slots.Next = mapIteratorNext
}

// ---- filter ----

type filterIterator interface {
	Object
	Callable() Object
	Iterator() Object
}

type filterIteratorStruct struct {
	Object
	callable Object
	iterator Object
}

func (fis *filterIteratorStruct) Callable() Object {
	return fis.callable
}

func (fis *filterIteratorStruct) Iterator() Object {
	return fis.iterator
}

var _ filterIterator = (*filterIteratorStruct)(nil)

func newFilterIterator(t Type, callable, iterator Object) filterIterator {
	return &filterIteratorStruct{newObject(t), callable, iterator}
}

// FilterIteratorType ..
var FilterIteratorType = newBuiltinType("filter")

func filterIteratorNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("filter", args, kwArgs, ObjectType, ObjectType)
	callable := args[0]
	iterator := Iter(args[1])
	return newFilterIterator(t, callable, iterator)
}

func filterIteratorIter(o Object) Object {
	return o
}

func filterIteratorNext(o Object) Object {
	fi := o.(filterIterator)
	callable := fi.Callable()
	iterator := fi.Iterator()
	for {
		val := NextDefault(iterator, AfterLast)
		if val == AfterLast {
			panic(Raise(StopIterationType))
		}
		var testval Object
		if callable == None {
			testval = val
		} else {
			testval = Cal(callable, val)
		}
		if IsTrue(testval) {
			return val
		}
	}
}

func initFilterIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.New = filterIteratorNew
	slots.Iter = filterIteratorIter
	slots.Next = filterIteratorNext
}

// ---- enumerate ----

type enumerateIterator interface {
	Object
	Iterator() Object
	Index() int
	increment()
}

type enumerateIteratorStruct struct {
	Object
	iterator Object
	index    int
}

func (eis *enumerateIteratorStruct) Iterator() Object {
	return eis.iterator
}

func (eis *enumerateIteratorStruct) Index() int {
	return eis.index
}

func (eis *enumerateIteratorStruct) increment() {
	eis.index++
}

var _ enumerateIterator = (*enumerateIteratorStruct)(nil)

func newEnumerateIterator(t Type, iterator Object, index int) enumerateIterator {
	return &enumerateIteratorStruct{newObject(t), iterator, index}
}

// EnumerateIteratorType ..
var EnumerateIteratorType = newBuiltinType("enumerate")

func enumerateIteratorNew(t Type, args Args, kwArgs KWArgs) Object {
	var index Object
	// TODO this arg checking is a mess
	if index = kwArgs.get("start", nil); index != nil {
		checkFunctionKWArgs("enumerate", kwArgs, map[string]Type{
			"start": IntType,
		})
		checkFunctionArgs("enumerate", args, nil, ObjectType)
	} else {
		checkFunctionArgsMin("enumerate", args, kwArgs, 1, ObjectType, IntType)
		if len(args) > 1 {
			index = args[1]
		} else {
			index = NewInt(0)
		}
	}
	iterator := Iter(args[0])
	return newEnumerateIterator(t, iterator, index.(Int).Value())
}

func enumerateIteratorIter(o Object) Object {
	return o
}

func enumerateIteratorNext(o Object) Object {
	ei := o.(enumerateIterator)
	val := NextDefault(ei.Iterator(), AfterLast)
	if val == AfterLast {
		panic(Raise(StopIterationType))
	}
	res := NewTuple(NewInt(ei.Index()), val)
	ei.increment()
	return res
}

func initEnumerateIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.New = enumerateIteratorNew
	slots.Iter = enumerateIteratorIter
	slots.Next = enumerateIteratorNext
}

// ---- reversed ----

type reversedIterator interface {
	Object
	Seq() Object
	Index() int
	Len() int
	setIndex(int)
}

type reversedIteratorStruct struct {
	Object
	seq   Object
	index int
	len   int
}

func (ris *reversedIteratorStruct) Seq() Object {
	return ris.seq
}

func (ris *reversedIteratorStruct) Index() int {
	return ris.index
}

func (ris *reversedIteratorStruct) Len() int {
	return ris.len
}

func (ris *reversedIteratorStruct) setIndex(index int) {
	ris.index = index
}

var _ reversedIterator = (*reversedIteratorStruct)(nil)

func newReversedIterator(seq Object, len int) reversedIterator {
	return &reversedIteratorStruct{newObject(ReversedIteratorType), seq, len - 1, len}
}

// ReversedIteratorType ..
var ReversedIteratorType = newBuiltinType("reversed")

func reversedIteratorNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("reversed", args, kwArgs, ObjectType)
	o := args[0]
	// try __reversed__
	if reversedFn := o.Type().Slots().Reversed; reversedFn != nil {
		return reversedFn(o)
	}
	// must be __len__ and __getitem__
	otype := o.Type()
	if otype.Slots().Len == nil || otype.Slots().GetItem == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf("'%s' object is not reversible", otype.Name())))
	}
	len := Len(o).Value()
	return newReversedIterator(o, len)
}

func reversedIteratorIter(o Object) Object {
	return o
}

func reversedIteratorNext(o Object) Object {
	ri := o.(reversedIterator)
	index := ri.Index()
	if index < 0 {
		// iteration already over
		panic(Raise(StopIterationType))
	}
	item := func() Object {
		defer Catch(IndexErrorType, func(ex BaseException) {
			// got IndexError, iteration over from now on
			index = 0
			panic(Raise(StopIterationType))
		}, nil)
		return GetItem(ri.Seq(), NewInt(index))
	}()
	ri.setIndex(index - 1)
	return item
}

func initReversedIteratorType(slots *typeSlots, dict map[string]Object) {
	slots.New = reversedIteratorNew
	slots.Iter = reversedIteratorIter
	slots.Next = reversedIteratorNext
}
