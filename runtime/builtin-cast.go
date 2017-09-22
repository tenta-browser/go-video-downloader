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
 * builtin-cast.go: Checked casts of interface{} to various types
 */

package runtime

import "fmt"

// CastToInt casts to int, handles boxing, panics with extractorError
func CastToInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return v
	case OptInt:
		return v.Get()
	default:
		panic(NewCastError(val, "int"))
	}
}

// CastToString casts to string, handles boxing, panics with extractorError
func CastToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case OptString:
		return v.Get()
	default:
		panic(NewCastError(val, "str"))
	}
}

// CastToOptInt casts to OptInt, handles boxing, panics with extractorError
func CastToOptInt(val interface{}) OptInt {
	switch v := val.(type) {
	case nil:
		return OptInt{}
	case int:
		return AsOptInt(v)
	case OptInt:
		return v
	default:
		panic(NewCastError(val, "OptInt"))
	}
}

// CastToOptString casts to OptString, handles boxing, panics with extractorError
func CastToOptString(val interface{}) OptString {
	switch v := val.(type) {
	case nil:
		return OptString{}
	case string:
		return AsOptString(v)
	case OptString:
		return v
	default:
		panic(NewCastError(val, "OptString"))
	}
}

// NewCastError creates extractorError's for failed casts
func NewCastError(val interface{}, destType string) error {
	return newExtractorError(fmt.Sprintf("Cannot cast %T to %s", val, destType))
}
