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
 * builtin-json.go: Re-implementation of parts of Python's json module
 */

package runtime

import "encoding/json"

// JSONLoads implements python/json.loads
func JSONLoads(s string) interface{} {
	res, err := jsonLoads(s)
	if err != nil {
		panic(newExtractorError("Failed to parse JSON: " + err.Error()))
	}
	return res
}

func jsonLoads(s string) (interface{}, error) {
	var res interface{}
	err := json.Unmarshal([]byte(s), &res)
	return res, err
}

// JSONDumps implements python/json.dumps
func JSONDumps(obj interface{}) string {
	res, err := jsonDumps(obj)
	if err != nil {
		panic(newExtractorError("Failed to dump JSON: " + err.Error()))
	}
	return res
}

func jsonDumps(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// MarshalJSON provides json marshalling for optional types
func (ostr OptString) MarshalJSON() ([]byte, error) {
	return json.Marshal(ostr.GetOrDef(""))
}

// MarshalJSON provides json marshalling for optional types
func (oint OptInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(oint.GetOrDef(0))
}

// MarshalJSON provides json marshalling for optional types
func (ofloat OptFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(ofloat.GetOrDef(0))
}
