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
 * method.go: Implementation of Python's method type
 */

package runtime

import "fmt"

// Method ..
type Method interface {
	Object
	Self() Object
	Fun() Object
}

type methodStruct struct {
	Object
	self Object
	fun  Object
}

func (ms *methodStruct) Self() Object {
	return ms.self
}

func (ms *methodStruct) Fun() Object {
	return ms.fun
}

var _ Method = (*methodStruct)(nil)

func newMethod(self Object, fun Object) Method {
	return &methodStruct{newObject(MethodType), self, fun}
}

// MethodType ..
var MethodType = newBuiltinType("method")

func methodNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("__new__", args, kwArgs, ObjectType, ObjectType)
	return newMethod(args[0], args[1])
}

func methodCall(o Object, args Args, kwArgs KWArgs) Object {
	m := o.(Method)
	methodArgs := append(Args{m.Self()}, args...)
	return Call(m.Fun(), methodArgs, kwArgs)
}

func methodRepr(o Object) Object {
	m := o.(Method)
	return NewStr(fmt.Sprintf("<bound method %s of %s>",
		Repr(m.Fun()).Value(), Repr(m.Self()).Value()))
}

func initMethodType(slots *typeSlots, dict map[string]Object) {
	slots.New = methodNew
	slots.Call = methodCall
	slots.Repr = methodRepr
}
