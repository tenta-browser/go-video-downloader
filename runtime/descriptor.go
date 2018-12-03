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
 * descriptor.go: Common descriptor implementations
 */

package runtime

// Property ..
type Property interface {
	Object
	Get() Object
	setGet(Object)
	Set() Object
	setSet(Object)
	Del() Object
	setDel(Object)
}

type propertyStruct struct {
	Object
	get Object
	set Object
	del Object
}

func (ps *propertyStruct) Get() Object {
	return ps.get
}

func (ps *propertyStruct) setGet(get Object) {
	ps.get = get
}

func (ps *propertyStruct) Set() Object {
	return ps.set
}

func (ps *propertyStruct) setSet(set Object) {
	ps.set = set
}

func (ps *propertyStruct) Del() Object {
	return ps.del
}

func (ps *propertyStruct) setDel(del Object) {
	ps.del = del
}

var _ Property = (*propertyStruct)(nil)

func newProperty(get, set, del Object) Property {
	return &propertyStruct{newObject(PropertyType), get, set, del}
}

// PropertyType ..
var PropertyType = newBuiltinType("property")

func propertyNew(t Type, args Args, kwArgs KWArgs) Object {
	return newProperty(nil, nil, nil)
}

func propertyInit(o Object, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("__init__", args, kwArgs, 0, ObjectType, ObjectType, ObjectType)
	p := o.(Property)
	argc := len(args)
	if argc > 0 {
		p.setGet(args[0])
	}
	if argc > 1 {
		p.setSet(args[1])
	}
	if argc > 2 {
		p.setDel(args[2])
	}
	return None
}

func propertyGet(desc Object, inst Object, t Type) Object {
	p := desc.(Property)
	getFn := p.Get()
	if getFn == nil || getFn == None {
		panic(RaiseType(AttributeErrorType, "unreadable attribute"))
	}
	return Cal(getFn, inst)
}

func propertySet(desc Object, inst Object, value Object) {
	p := desc.(Property)
	setFn := p.Set()
	if setFn == nil || setFn == None {
		panic(RaiseType(AttributeErrorType, "can't set attribute"))
	}
	Cal(setFn, inst, value)
}

func propertyDelete(desc Object, inst Object) {
	p := desc.(Property)
	delFn := p.Del()
	if delFn == nil || delFn == None {
		panic(RaiseType(AttributeErrorType, "can't delete attribute"))
	}
	Cal(delFn, inst)
}

func initPropertyType(slots *typeSlots, dict map[string]Object) {
	slots.New = propertyNew
	slots.Init = propertyInit
	slots.Get = propertyGet
	slots.Set = propertySet
	slots.Delete = propertyDelete
}
