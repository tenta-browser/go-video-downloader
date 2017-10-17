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
 * builtin-cast.go: Conversion utilities from interface{} types to various types
 */

package runtime

import (
	"fmt"
	"strconv"
)

// ConvertToInt implements python/int()
func ConvertToInt(val interface{}) int {
	atoi := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(newExtractorError(fmt.Sprintf("Failed to convert to int: '%s'", s)))
		}
		return i
	}
	switch v := val.(type) {
	case int:
		return v
	case OptInt:
		return v.Get()
	case string:
		return atoi(v)
	case OptString:
		return atoi(v.Get())
	default:
		panic(newConvertError(val, "int"))
	}
}

// ConvertToString implements python/str()
func ConvertToString(val interface{}) string {
	switch v := val.(type) {
	case int:
		return string(v)
	case OptInt:
		return string(v.Get())
	case string:
		return v
	case OptString:
		return v.Get()
	default:
		panic(newConvertError(val, "string"))
	}
}

func newConvertError(val interface{}, destType string) error {
	return newExtractorError(fmt.Sprintf("Cannot convert %T to %s", val, destType))
}
