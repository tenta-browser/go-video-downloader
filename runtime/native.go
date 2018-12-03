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
 * native.go: Implementation of a native type, for passing Go objects into
 *            Python land and back
 */

package runtime

import "fmt"

// Native objects wrap a Go object
type Native interface {
	Object
	Value() interface{}
}

type nativeStruct struct {
	Object
	value interface{}
}

func (ns *nativeStruct) Value() interface{} {
	return ns.value
}

var _ Native = (*nativeStruct)(nil)

// NewNative wraps value into a Native
func NewNative(value interface{}) Native {
	return &nativeStruct{newObject(NativeType), value}
}

// NativeType is the type object of Native objects
var NativeType = newBuiltinType("native")

func nativeRepr(o Object) Object {
	n := o.(Native)
	value := n.Value()
	return NewStr(fmt.Sprintf("<native object %T %v>", value, value))
}

func initNativeType(slots *typeSlots, dict map[string]Object) {
	slots.Repr = nativeRepr
}
