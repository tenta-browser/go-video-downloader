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
 * builtin.go: Re-implementation of various builtin functions of Python
 */

package runtime

const (
	// OperatorConcat represents python/operator.concat
	OperatorConcat = 1
)

// Todo is used as a replacement for functions not yet implemented
func Todo() {
	panic(newExtractorError("ToOodoOo"))
}

// UnsafeSubscript tries to subscript val no matter if it's a list or a map
func UnsafeSubscript(val interface{}, idx interface{}) interface{} {
	switch col := val.(type) {
	case []interface{}:
		return col[idx.(int)]
	case map[string]interface{}:
		return col[idx.(string)]
	default:
		panic(newExtractorError("Tried to subscript non-collection type"))
	}
}

// DictGet implements python/dict.get
func DictGet(dict map[string]interface{}, key string, def interface{}) interface{} {
	if val, found := dict[key]; found {
		return val
	}
	return def
}

// DictContains checks if a key is present in a dictionary
func DictContains(dict map[string]interface{}, key string) bool {
	_, found := dict[key]
	return found
}

// DictUpdate implements python/dict.update
func DictUpdate(dict map[string]interface{}, other map[string]interface{}) {
	for k, v := range other {
		dict[k] = v
	}
}

// IsTruthy tells whether the supplied interface value holds a truthy value
// (non-empty string/list, True boolean, >0 int, etc.)
func IsTruthy(val interface{}) bool {
	switch v := val.(type) {
	case OptInt:
		return v.IsSet()
	case OptString:
		return v.IsSet()
	case int:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v != ""
	case bool:
		return v
	default:
		return v != nil
	}
}

// FuncReduce implements functools.reduce
func FuncReduce(op int, l []interface{}) interface{} {
	if op == OperatorConcat {
		result := ""
		for _, el := range l {
			result += CastToString(el)
		}
		return result
	}
	panic(newExtractorError("Unknown operator: " + string(op)))
}
