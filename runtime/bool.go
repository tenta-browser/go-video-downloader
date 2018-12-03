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
 * bool.go: Implementation of Python's boolean type
 */

package runtime

// NewBool ..
func NewBool(value bool) Int {
	if value {
		return True
	}
	return False
}

// BoolType ..
var BoolType = newBuiltinType("bool", IntType)

func boolNew(t Type, args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("bool", args, kwArgs, 0, ObjectType)
	if len(args) == 0 {
		return False
	}
	return NewBool(IsTrue(args[0]))
}

func boolRepr(o Object) Object {
	i := o.(Int)
	if i.Value() == 0 {
		return falseStr
	}
	return trueStr
}

func initBoolType(slots *typeSlots, dict map[string]Object) {
	slots.New = boolNew
	slots.Repr = boolRepr
}

var (
	// True ..
	True = newInt(BoolType, 1)
	// False ..
	False    = newInt(BoolType, 0)
	trueStr  = NewStr("True")
	falseStr = NewStr("False")
)
