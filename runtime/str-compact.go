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
 * str-compact.go: Compact implementation variant of the string type
 */

package runtime

// StrCompact ..
type StrCompact string

var strCompactBase Object

// Type ..
func (sc StrCompact) Type() Type {
	return strCompactBase.Type()
}

// Dict ..
func (sc StrCompact) Dict() Dict {
	return strCompactBase.Dict()
}

func (sc StrCompact) setDict(d Dict) {
	strCompactBase.setDict(d)
}

// IsInstance ..
func (sc StrCompact) IsInstance(t Type) bool {
	return strCompactBase.IsInstance(t)
}

// Value ..
func (sc StrCompact) Value() string {
	return string(sc)
}

func newStrCompact(value string) Str {
	if strCompactBase == nil {
		strCompactBase = newObject(StrType)
	}
	return StrCompact(value)
}
