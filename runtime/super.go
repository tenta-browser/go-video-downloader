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
 * super.go: Implementation of Python's super function
 */

package runtime

// Super ..
type Super interface {
	Object
	Sub() Type
	Obj() Object
	ObjType() Type
}

type superStruct struct {
	Object
	sub     Type
	obj     Object
	objType Type
}

func (ss *superStruct) Sub() Type {
	return ss.sub
}

func (ss *superStruct) Obj() Object {
	return ss.obj
}

func (ss *superStruct) ObjType() Type {
	return ss.objType
}

var _ Super = (*superStruct)(nil)

func newSuper(t Type, sub Type, obj Object, objType Type) Super {
	return &superStruct{newObject(t), sub, obj, objType}
}

// SuperType ..
var SuperType = newBuiltinType("super")

func superNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("__init__", args, kwArgs, TypeType, ObjectType)
	sub := args[0].(Type)
	var obj Object
	var objType Type
	if args[1].IsInstance(TypeType) && args[1].(Type).IsSubType(sub) {
		objType = args[1].(Type)
	} else if args[1].IsInstance(sub) {
		obj = args[1]
		objType = obj.Type()
	} else {
		panic(RaiseType(TypeErrorType,
			"super(type, obj): obj must be an instance or subtype of type"))
	}
	return newSuper(t, sub, obj, objType)
}

func superGetAttribute(o Object, name Str) Object {
	sup := o.(Super)
	// Tell the truth about the __class__ attribute.
	if sup.ObjType() != nil && name.Value() != "__class__" {
		mro := sup.ObjType().Mro()
		n := len(mro)
		// Start from the immediate mro successor to the specified type.
		i := 0
		for i < n && mro[i] != sup.Sub() {
			i++
		}
		i++
		var inst Object
		if sup.Obj() != nil {
			inst = sup.Obj()
		}
		// Now do normal mro lookup from the successor type.
		for ; i < n; i++ {
			dict := mro[i].Dict()
			res := dict.GetItem(name)
			if res == nil {
				continue
			}
			usedSymbolAttr(mro[i], name)
			if getFn := res.Type().Slots().Get; getFn != nil {
				// Found a descriptor so invoke it.
				return getFn(res, inst, sup.ObjType())
			}
			return res
		}
	}
	// Attribute not found on base classes so lookup the attr on the super
	// object itself. Most likely will AttributeError.
	return objectGetAttribute(o, name)
}

func initSuperType(slots *typeSlots, dict map[string]Object) {
	slots.New = superNew
	slots.GetAttribute = superGetAttribute
}
