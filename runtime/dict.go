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
 * dict.go: Implementation of Python's dictionary type
 */

package runtime

import (
	"fmt"
	"strings"
)

// Dict ..
type Dict interface {
	Object

	GetItem(key Object) Object
	SetItem(key Object, value Object)
	DelItem(key Object) Object
	SetDefault(key Object, value Object) Object
	Len() int
	Iterator() hashtableIterator
	Keys() []Object

	Update(iterable Object)
}

type dictStruct struct {
	Object
	table hashtable
}

func (ds *dictStruct) GetItem(key Object) Object {
	return ds.table.get(key)
}

func (ds *dictStruct) SetItem(key Object, value Object) {
	ds.table.put(key, value, true)
}

func (ds *dictStruct) DelItem(key Object) Object {
	return ds.table.remove(key)
}

func (ds *dictStruct) SetDefault(key, value Object) Object {
	oldValue := ds.table.put(key, value, false)
	if oldValue == nil {
		oldValue = value
	}
	return oldValue
}

func (ds *dictStruct) Len() int {
	return ds.table.len
}

func (ds *dictStruct) Iterator() hashtableIterator {
	return ds.table.iterator()
}

func (ds *dictStruct) Keys() []Object {
	keys := make([]Object, ds.Len())
	it := ds.Iterator()
	i := 0
	for e := it(); e != nil; e = it() {
		keys[i] = e.Key
		i++
	}
	return keys
}

func (ds *dictStruct) Update(o Object) {
	if o.IsInstance(DictType) {
		o = dictItemList(o)
	}
	i := 0
	seqForEach(o, func(item Object) {
		seqApply(item, func(keyValue []Object) {
			if len := len(keyValue); len != 2 {
				panic(RaiseType(ValueErrorType, fmt.Sprintf(
					"dictionary update sequence element #%d has length %d; 2 is required",
					i, len)))
			}
			ds.SetItem(keyValue[0], keyValue[1])
		})
		i++
	})
}

var _ Dict = (*dictStruct)(nil)

func newDict(t Type) Dict {
	return &dictStruct{newObject(t), *newHashTable(4)}
}

// NewDictWithTable ..
func NewDictWithTable(table map[Object]Object) Dict {
	d := newDict(DictType)
	for key, value := range table {
		d.SetItem(key, value)
	}
	return d
}

// NewDict ..
func NewDict() Dict {
	return newDict(DictType)
}

// DictType ..
var DictType = newBuiltinType("dict")

func dictNew(t Type, args Args, kwArgs KWArgs) Object {
	return newDict(t)
}

func dictInit(o Object, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("dict", args, nil, 0, ObjectType)
	d := o.(Dict)
	if len(args) > 0 {
		d.Update(args[0])
	}
	for _, kwArg := range kwArgs {
		d.SetItem(NewStr(kwArg.Name), kwArg.Value)
	}
	return None
}

func dictRepr(o Object) Object {
	b := strings.Builder{}
	iterator := o.(Dict).Iterator()
	b.WriteByte('{')
	i := 0
	for e := iterator(); e != nil; e = iterator() {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(Repr(e.Key).Value())
		b.WriteString(": ")
		b.WriteString(Repr(e.Value).Value())
		i++
	}
	b.WriteByte('}')
	return NewStr(b.String())
}

func dictGetItem(o Object, key Object) Object {
	d := o.(Dict)
	value := d.GetItem(key)
	if value == nil {
		panic(RaiseType(KeyErrorType, Repr(key).Value()))
	}
	return value
}

func dictSetItem(o Object, key Object, value Object) {
	d := o.(Dict)
	d.SetItem(key, value)
}

func dictDelItem(o Object, key Object) {
	d := o.(Dict)
	value := d.DelItem(key)
	if value == nil {
		panic(RaiseType(KeyErrorType, Repr(key).Value()))
	}
}

func dictLen(o Object) Object {
	d := o.(Dict)
	return NewInt(d.Len())
}

func dictIter(o Object) Object {
	// TODO proper key iterator
	return Iter(NewList(o.(Dict).Keys()...))
}

func dictValueList(o Object) Object {
	d := o.(Dict)
	it := d.Iterator()
	res := []Object{}
	for e := it(); e != nil; e = it() {
		res = append(res, e.Value)
	}
	return NewList(res...)
}

func dictItemList(o Object) Object {
	d := o.(Dict)
	it := d.Iterator()
	res := []Object{}
	for e := it(); e != nil; e = it() {
		res = append(res, NewTuple(e.Key, e.Value))
	}
	return NewList(res...)
}

func dictsAreEqual(d1, d2 Dict) bool {
	if d1 == d2 {
		return true
	}
	if d1.Len() != d2.Len() {
		return false
	}
	iter1 := d1.Iterator()
	for el1 := iter1(); el1 != nil; el1 = iter1() {
		key := el1.Key
		value1 := el1.Value
		value2 := d2.GetItem(key)
		if value2 == nil || !IsTrue(Eq(value1, value2)) {
			return false
		}
	}
	return true
}

func dictEq(a, b Object) Object {
	da := a.(Dict)
	if !b.IsInstance(DictType) {
		return NotImplemented
	}
	db := b.(Dict)
	return NewBool(dictsAreEqual(da, db))
}

func dictNe(a, b Object) Object {
	da := a.(Dict)
	if !b.IsInstance(DictType) {
		return NotImplemented
	}
	db := b.(Dict)
	return NewBool(!dictsAreEqual(da, db))
}

func dictContains(o Object, needle Object) Object {
	d := o.(Dict)
	return NewBool(d.GetItem(needle) != nil)
}

func dictCopy(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("copy", args, kwArgs, DictType)
	return Call(DictType, args, kwArgs)
}

func dictGet(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("get", args, kwArgs, 2, DictType, ObjectType, ObjectType)
	d := args[0].(Dict)
	key := args[1]
	val := d.GetItem(key)
	if val == nil {
		if len(args) > 2 {
			return args[2]
		}
		return None
	}
	return val
}

func dictItems(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("items", args, kwArgs, DictType)
	// TODO proper iterator here, eliminate list functions
	return Iter(dictItemList(args[0]))
}

func dictKeys(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("keys", args, kwArgs, DictType)
	// TODO proper iterator here, eliminate list functions
	return Iter(NewList(args[0].(Dict).Keys()...))
}

func dictUpdate(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("update", args, nil, 1, DictType, ObjectType)
	d := args[0].(Dict)
	if len(args) > 1 {
		d.Update(args[1])
	}
	for _, kwArg := range kwArgs {
		d.SetItem(NewStr(kwArg.Name), kwArg.Value)
	}
	return None
}

func dictSetDefault(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("setdefault", args, kwArgs, 2,
		DictType, ObjectType, ObjectType)
	d := args[0].(Dict)
	key := args[1]
	def := None
	if len(args) > 2 {
		def = args[2]
	}
	return d.SetDefault(key, def)
}

func dictValues(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("values", args, kwArgs, DictType)
	// TODO proper iterator here, eliminate list functions
	return Iter(dictValueList(args[0]))
}

func initDictType(slots *typeSlots, dict map[string]Object) {
	dict["copy"] = newBuiltinFunction("copy", dictCopy)
	dict["get"] = newBuiltinFunction("get", dictGet)
	dict["items"] = newBuiltinFunction("items", dictItems)
	dict["keys"] = newBuiltinFunction("keys", dictKeys)
	dict["update"] = newBuiltinFunction("update", dictUpdate)
	dict["setdefault"] = newBuiltinFunction("setdefault", dictSetDefault)
	dict["values"] = newBuiltinFunction("values", dictValues)
	slots.New = dictNew
	slots.Init = dictInit
	slots.Hash = unhashableHash
	slots.Repr = dictRepr
	slots.GetItem = dictGetItem
	slots.SetItem = dictSetItem
	slots.DelItem = dictDelItem
	slots.Len = dictLen
	slots.Iter = dictIter
	slots.Eq = dictEq
	slots.Ne = dictNe
	slots.Contains = dictContains
}

// ---- Simplistic hashtable implementation ----

type hashtable struct {
	len   int
	table []*hashtableEntry
}

type hashtableEntry struct {
	Key   Object
	Value Object
	next  *hashtableEntry
}

type hashtableIterator func() *hashtableEntry

func newHashTable(capacity int) *hashtable {
	return &hashtable{len: 0, table: make([]*hashtableEntry, capacity)}
}

func (ht *hashtable) slot(key Object) uint {
	h := Hash(key).Value()
	return uint(h) % uint(len(ht.table))
}

func (ht *hashtable) get(key Object) Object {
	s := ht.slot(key)
	for e := ht.table[s]; e != nil; e = e.next {
		if IsTrue(Eq(key, e.Key)) {
			return e.Value
		}
	}
	return nil
}

func (ht *hashtable) put(key, value Object, overwrite bool) Object {
	s := ht.slot(key)
	pe := &ht.table[s]
	for {
		e := *pe
		if e == nil {
			// reached end, add new entry
			*pe = &hashtableEntry{
				Key:   key,
				Value: value,
			}
			ht.len++
			if ht.len*2 > len(ht.table) {
				// 50% load reached, grow table
				ht.grow()
			}
			return nil
		}
		if IsTrue(Eq(key, e.Key)) {
			// found, replace entry
			oldValue := e.Value
			if overwrite {
				e.Value = value
			}
			return oldValue
		}
		pe = &e.next
	}
}

func (ht *hashtable) grow() {
	oldTable := ht.table
	ht.len = 0
	ht.table = make([]*hashtableEntry, 2*len(oldTable))
	for _, e := range oldTable {
		for ; e != nil; e = e.next {
			ht.put(e.Key, e.Value, false)
		}
	}
}

func (ht *hashtable) remove(key Object) Object {
	s := ht.slot(key)
	pe := &ht.table[s]
	for {
		e := *pe
		if e == nil {
			// reached end, nothing to remove
			return nil
		}
		if IsTrue(Eq(key, e.Key)) {
			// found, unlink entry
			*pe = e.next
			return e.Value
		}
		pe = &e.next
	}
}

func (ht *hashtable) iterator() hashtableIterator {
	s := -1
	var e *hashtableEntry
	return func() *hashtableEntry {
		for e == nil {
			// advance to next slot
			s++
			if s >= len(ht.table) {
				return nil
			}
			e = ht.table[s]
		}
		r := e
		e = e.next
		return r
	}
}
