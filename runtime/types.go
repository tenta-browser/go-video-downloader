/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 * types.py: Custom types needed to simulate Python's types
 */

package runtime

// OptString represents an optional string
type OptString struct {
	string
	set bool
}

// AsOptString converts a string to an optional string
func AsOptString(value string) OptString {
	return OptString{value, true}
}

// Get returns the value of the optional object
func (obj OptString) Get() string {
	return obj.string
}

// GetOrDef returns the value of the optional object is it's available
// or the default value otherwise
func (obj OptString) GetOrDef(def string) string {
	if !obj.IsSet() {
		return def
	}
	return obj.string
}

// Set sets the value of the optional object
func (obj *OptString) Set(value string) {
	obj.string = value
	obj.set = true
}

// IsSet returns whether the optional object has a value
func (obj OptString) IsSet() bool {
	return obj.set
}

// OptInt represents an optional int
type OptInt struct {
	int
	set bool
}

// AsOptInt converts an int to an optional int
func AsOptInt(value int) OptInt {
	return OptInt{value, true}
}

// Get returns the value of the optional object
func (obj OptInt) Get() int {
	return obj.int
}

// GetOrDef returns the value of the optional object is it's available
// or the default value otherwise
func (obj OptInt) GetOrDef(def int) int {
	if !obj.IsSet() {
		return def
	}
	return obj.int
}

// Set sets the value of the optional object
func (obj *OptInt) Set(value int) {
	obj.int = value
	obj.set = true
}

// IsSet returns whether the optional object has a value
func (obj OptInt) IsSet() bool {
	return obj.set
}

type noDefaultType struct{}

// NoDefault is the simulation of youtube-dl's hack NO_DEFAULT = object() (utils.py),
// which is used in some places as a marker value (where None meant something else)
var NoDefault = noDefaultType{}

func isNoDefault(something interface{}) bool {
	_, ok := something.(noDefaultType)
	return ok
}
