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
 * type.go: Implementation of Python's base type
 */

package runtime

import (
	"fmt"
	"reflect"
)

// Type ..
type Type interface {
	Object
	Name() string
	Bases() []Type
	Mro() []Type
	setMro(mro []Type)
	Slots() *typeSlots
	Builtin() bool

	IsSubType(super Type) bool
}

type typeStruct struct {
	Object
	name    string
	bases   []Type
	mro     []Type
	slots   typeSlots
	builtin bool
}

func (ts *typeStruct) Name() string {
	return ts.name
}

func (ts *typeStruct) Bases() []Type {
	return ts.bases
}

func (ts *typeStruct) Mro() []Type {
	return ts.mro
}

func (ts *typeStruct) setMro(mro []Type) {
	ts.mro = mro
}

func (ts *typeStruct) Slots() *typeSlots {
	return &ts.slots
}

func (ts *typeStruct) Builtin() bool {
	return ts.builtin
}

func (ts typeStruct) IsSubType(super Type) bool {
	for _, b := range ts.mro {
		if b == super {
			return true
		}
	}
	return false
}

var _ Type = (*typeStruct)(nil)

func newType(name string, bases []Type, builtin bool) Type {
	if len(bases) == 0 {
		bases = append(bases, ObjectType)
	}
	return &typeStruct{
		Object:  &objectStruct{typ: TypeType},
		name:    name,
		bases:   bases,
		builtin: builtin,
	}
}

func newBuiltinType(name string, bases ...Type) Type {
	return newType(name, bases, true)
}

func fillDictFromSlots(t Type, dict map[string]Object) {
	for name, function := range t.Slots().makeFunctions(t) {
		dict[name] = function
	}
}

func postInitType(t Type) {
	// calculate mro
	mro := mroCalc(t)
	if mro == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf("mro error for: %s", t.Name())))
	}
	t.setMro(mro)
	// inherit slots
	slotsValue := reflect.ValueOf(t.Slots()).Elem()
	for i := 0; i < slotCount; i++ {
		slotField := slotsValue.Field(i)
		if !slotField.IsNil() {
			continue
		}
		for _, base := range mro {
			baseSlotFn := reflect.ValueOf(base.Slots()).Elem().Field(i)
			if !baseSlotFn.IsNil() {
				slotField.Set(baseSlotFn)
				break
			}
		}
	}
}

func mroMerge(seqs [][]Type) []Type {
	var res []Type
	numSeqs := len(seqs)
	hasNonEmptySeqs := true
	for hasNonEmptySeqs {
		var cand Type
		for i := 0; i < numSeqs && cand == nil; i++ {
			// The next candidate will be absent from or at the head
			// of all lists. If we try a candidate and we find it's
			// somewhere past the head of one of the lists, reject.
			seq := seqs[i]
			if len(seq) == 0 {
				continue
			}
			cand = seq[0]
		RejectCandidate:
			for _, seq := range seqs {
				numElems := len(seq)
				for j := 1; j < numElems; j++ {
					if seq[j] == cand {
						cand = nil
						break RejectCandidate
					}
				}
			}
		}
		if cand == nil {
			// We could not find a candidate meaning that the
			// hierarchy is inconsistent.
			return nil
		}
		res = append(res, cand)
		hasNonEmptySeqs = false
		for i, seq := range seqs {
			if len(seq) > 0 {
				if seq[0] == cand {
					// Remove the candidate from all lists
					// (it will only be found at the head of
					// any list because otherwise it would
					// have been rejected above.)
					seqs[i] = seq[1:]
				}
				if len(seqs[i]) > 0 {
					hasNonEmptySeqs = true
				}
			}
		}
	}
	return res
}

func mroCalc(t Type) []Type {
	seqs := [][]Type{{t}}
	for _, b := range t.Bases() {
		seqs = append(seqs, b.Mro())
	}
	seqs = append(seqs, t.Bases())
	return mroMerge(seqs)
}

func mroLookup(t Type, name Str) Object {
	for _, t := range t.Mro() {
		if d := t.Dict(); d != nil {
			if v := t.Dict().GetItem(name); v != nil {
				usedSymbolAttr(t, name)
				return v
			}
		}
	}
	return nil
}

func newClass(name string, bases []Type, dict Dict) Type {
	t := newType(name, bases, false)
	t.Slots().fillSlots(dict)
	postInitType(t)
	t.setDict(dict)
	return t
}

// TypeType ..
var TypeType = &typeStruct{
	name:    "type",
	bases:   []Type{ObjectType},
	builtin: true,
}

func typeNew(t Type, args Args, kwArgs KWArgs) Object {
	switch len(args) {
	case 1:
		// type(obj) returns obj's type
		return args[0].Type()
	case 3:
		break // continued below
	default:
		panic(RaiseType(TypeErrorType, "type() takes 1 or 3 arguments"))
	}
	checkFunctionArgs("type", args, kwArgs, StrType, TupleType, DictType)
	name := args[0].(Str).Value()
	bases := args[1].(Tuple).Elems()
	dict := args[2].(Dict)
	baseTypes := make([]Type, len(bases))
	for i, base := range bases {
		if !base.IsInstance(TypeType) {
			panic(RaiseType(TypeErrorType,
				fmt.Sprintf("not a valid base class: %s", Repr(base).Value())))
		}
		baseTypes[i] = base.(Type)
	}
	return newClass(name, baseTypes, dict)
}

func typeCall(to Object, args Args, kwArgs KWArgs) Object {
	t := to.(Type)
	newFn := t.Slots().New
	if newFn == nil {
		panic(RaiseType(TypeErrorType, fmt.Sprintf("type %s has no __new__", t.Name())))
	}
	o := newFn(t, args, kwArgs)
	if initFn := o.Type().Slots().Init; initFn != nil {
		initFn(o, args, kwArgs)
	}
	return o
}

// mostly copied from objectGetAttribute -- instance dictionary lookup changed
func typeGetAttribute(o Object, name Str) Object {
	// Data descriptors in type (class dictionary)
	var typeGetFn func(Object, Object, Type) Object
	typeAttr := mroLookup(o.Type(), name)
	if typeAttr != nil {
		typeAttrSlots := typeAttr.Type().Slots()
		typeGetFn = typeAttr.Type().Slots().Get
		if typeGetFn != nil && (typeAttrSlots.Set != nil || typeAttrSlots.Delete != nil) {
			return typeGetFn(typeAttr, o, o.Type())
		}
	}
	// Object in object's dict (instance dictionary)
	if attr := mroLookup(o.(Type), name); attr != nil {
		if getFn := attr.Type().Slots().Get; getFn != nil {
			return getFn(attr, None, o.(Type))
		}
		return attr
	}
	// Non-data descriptors in type (class dictionary)
	if typeGetFn != nil {
		return typeGetFn(typeAttr, o, o.Type())
	}
	// Object in type (class dictionary)
	if typeAttr != nil {
		return typeAttr
	}
	// give up
	panic(RaiseType(AttributeErrorType, fmt.Sprintf("type object '%s' has no attribute '%s'",
		o.(Type).Name(), name.Value())))
}

func typeRepr(o Object) Object {
	t := o.(Type)
	return NewStr(fmt.Sprintf("<class '%s'>", t.Name()))
}

func typeGetName(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("_get_name", args, kwArgs, TypeType)
	return NewStr(args[0].(Type).Name())
}

func initTypeType(slots *typeSlots, dict map[string]Object) {
	TypeType.Object = newObject(TypeType)
	dict["__name__"] = newProperty(
		newBuiltinFunction("_get_name", typeGetName),
		nil, nil)
	slots.New = typeNew
	slots.Call = typeCall
	slots.GetAttribute = typeGetAttribute
	slots.Repr = typeRepr
}

// ---- Class dict literal ----

// ClassDictLiteral ..
func ClassDictLiteral(m interface{}) Dict {
	return dictLiteral(m, true)
}
