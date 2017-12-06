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
	i, err := convertToInt(val)
	if err != nil {
		panic(newExtractorError(fmt.Sprintf("Failed to convert to int: '%v'", val)))
	}
	return i
}

func convertToInt(val interface{}) (int, error) {
	switch v := val.(type) {
	case int:
		return v, nil
	case OptInt:
		return v.Get(), nil
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return i, nil
	case OptString:
		i, err := strconv.Atoi(v.Get())
		if err != nil {
			return 0, err
		}
		return i, nil
	case float64:
		return int(v), nil
	case OptFloat:
		return int(v.Get()), nil
	default:
		panic(newConvertError(val, "int"))
	}
}

// ConvertToString implements python/str()
func ConvertToString(val interface{}) string {
	itoa := strconv.Itoa
	ftoa := func(f float64) string {
		return fmt.Sprintf("%v", f)
	}
	switch v := val.(type) {
	case int:
		return itoa(v)
	case OptInt:
		return itoa(v.Get())
	case string:
		return v
	case OptString:
		return v.Get()
	case float64:
		return ftoa(v)
	case OptFloat:
		return ftoa(v.Get())
	default:
		panic(newConvertError(val, "string"))
	}
}

// ConvertToFloat implements python/float()
func ConvertToFloat(val interface{}) float64 {
	f, err := convertToFloat(val)
	if err != nil {
		panic(newExtractorError(fmt.Sprintf("Failed to convert to float: '%v'", val)))
	}
	return f
}

func convertToFloat(val interface{}) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case OptInt:
		return float64(v.Get()), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	case OptString:
		f, err := strconv.ParseFloat(v.Get(), 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	case float64:
		return v, nil
	case OptFloat:
		return v.Get(), nil
	default:
		panic(newConvertError(val, "float"))
	}
}

// ConvertToList implements python/list()
func ConvertToList(val interface{}) []interface{} {
	// handled specially by the transpiler
	panic(newConvertError(val, "list"))
}

// ConvertToDict implements python/dict()
func ConvertToDict(val interface{}) SDict {
	// handled specially by the transpiler
	panic(newConvertError(val, "dict"))
}

// ConvertToSet implements python/set()
func ConvertToSet(val interface{}) map[string]struct{} {
	// handled specially by the transpiler
	panic(newConvertError(val, "set"))
}

func newConvertError(val interface{}, destType string) error {
	return newKindedExtractorError("convert", fmt.Sprintf("Cannot convert %T to %s", val, destType))
}
