/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
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
 * int-compact.go: Compact implementation variant of the integer type
 */

package runtime

// IntCompact ..
type IntCompact int

var intCompactBase Object

// Type ..
func (ic IntCompact) Type() Type {
	return intCompactBase.Type()
}

// Dict ..
func (ic IntCompact) Dict() Dict {
	return intCompactBase.Dict()
}

func (ic IntCompact) setDict(d Dict) {
	intCompactBase.setDict(d)
}

// IsInstance ..
func (ic IntCompact) IsInstance(t Type) bool {
	return intCompactBase.IsInstance(t)
}

// Value ..
func (ic IntCompact) Value() int {
	return int(ic)
}

func newIntCompact(value int) Int {
	if intCompactBase == nil {
		intCompactBase = newObject(IntType)
	}
	return IntCompact(value)
}
