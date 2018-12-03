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
 * object.go: Implementation of Python's base object type
 */

package runtime

import (
	"fmt"
	"reflect"
)

// Object ..
type Object interface {
	Type() Type
	Dict() Dict
	setDict(Dict)

	IsInstance(t Type) bool
}

type objectStruct struct {
	typ  Type
	dict Dict
}

func (os *objectStruct) Type() Type {
	return os.typ
}

func (os *objectStruct) Dict() Dict {
	return os.dict
}

func (os *objectStruct) setDict(d Dict) {
	os.dict = d
}

func (os *objectStruct) IsInstance(t Type) bool {
	return os.typ.IsSubType(t)
}

var _ Object = (*objectStruct)(nil)

func newObject(t Type) Object {
	o := &objectStruct{typ: t}
	if !t.Builtin() {
		o.setDict(NewDict())
	}
	return o
}

// ObjectType ..
var ObjectType = &typeStruct{
	name:    "object",
	builtin: true,
}

func objectNew(t Type, args Args, kwArgs KWArgs) Object {
	return newObject(t)
}

func objectGetAttribute(o Object, name Str) Object {
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
	if d := o.Dict(); d != nil {
		if value := d.GetItem(name); value != nil {
			return value
		}
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
	panic(RaiseType(AttributeErrorType,
		fmt.Sprintf(errGetAttribute, o.Type().Name(), name.Value())))
}

func objectSetAttr(o Object, name Str, value Object) {
	// data descriptor
	if typeAttr := mroLookup(o.Type(), name); typeAttr != nil {
		if typeSetFn := typeAttr.Type().Slots().Set; typeSetFn != nil {
			typeSetFn(typeAttr, o, value)
			return
		}
	}
	// value in object's dict
	if d := o.Dict(); d != nil {
		d.SetItem(name, value)
		return
	}
	// give up
	panic(RaiseType(AttributeErrorType,
		fmt.Sprintf(errGetAttribute, o.Type().Name(), name.Value())))
}

func objectDelAttr(o Object, name Str) {
	// data descriptor
	if typeAttr := mroLookup(o.Type(), name); typeAttr != nil {
		if typeDelFn := typeAttr.Type().Slots().Delete; typeDelFn != nil {
			typeDelFn(typeAttr, o)
			return
		}
	}
	// value in object's dict
	if d := o.Dict(); d != nil {
		if d.DelItem(name) != nil {
			return
		}
	}
	// give up
	panic(RaiseType(AttributeErrorType,
		fmt.Sprintf(errGetAttribute, o.Type().Name(), name.Value())))
}

func objectHash(o Object) Object {
	return NewInt(int(reflect.ValueOf(o).Pointer()))
}

func objectNe(a, b Object) Object {
	eq := tryCompare(compareEq, a, b)
	return NewBool(eq == NotImplemented || !IsTrue(eq))
}

func objectGetDict(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("_get_dict", args, kwArgs, ObjectType)
	o := args[0]
	d := o.Dict()
	if d == nil {
		panic(RaiseType(AttributeErrorType,
			fmt.Sprintf(errGetAttribute, o.Type().Name(), "__dict__")))
	}
	return d
}

func objectSetDict(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("_set_dict", args, kwArgs, ObjectType, DictType)
	args[0].setDict(args[1].(Dict))
	return None
}

func objectGetClass(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("_get_class", args, kwArgs, ObjectType)
	return args[0].Type()
}

func initObjectType(slots *typeSlots, dict map[string]Object) {
	ObjectType.Object = newObject(TypeType)
	dict["__dict__"] = newProperty(
		newBuiltinFunction("_get_dict", objectGetDict),
		newBuiltinFunction("_set_dict", objectSetDict),
		nil)
	dict["__class__"] = newProperty(
		newBuiltinFunction("_get_class", objectGetClass),
		nil, nil,
	)
	slots.New = objectNew
	slots.GetAttribute = objectGetAttribute
	slots.SetAttr = objectSetAttr
	slots.DelAttr = objectDelAttr
	slots.Hash = objectHash
	slots.Ne = objectNe
}
