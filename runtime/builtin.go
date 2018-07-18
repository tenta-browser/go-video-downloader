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

import (
	"fmt"
	"reflect"
)

const (
	// OperatorConcat represents python/operator.concat
	OperatorConcat = 1
)

// InitSliceArgs prepares the supplied (optional) slice arguments so that they can be used
// in for loops directly to calculate the resulting slice:
//  - ascending:  for i := low; i < high; i += step { .. }
//  - descending: for i := high; i > low; i += step { .. }
func InitSliceArgs(len int, low, high, step OptInt) (int, int, int) {
	var lo, hi int
	var min, max int
	// default step to 1; panic on zero step
	st := step.GetOrDef(1)
	if st == 0 {
		panic(newExtractorError("Slicing with step zero"))
	} else if st > 0 {
		min, max = 0, len
	} else {
		min, max = -1, len-1
		low, high = high, low
	}
	// default low to 0 if ascending, -1 if descending; handle negative index
	if low.IsSet() {
		lo = low.Get()
		if lo < 0 {
			lo += len
		}
	} else {
		lo = min
	}
	// default high to len if ascending, len-1 if descending; handle negative index
	if high.IsSet() {
		hi = high.Get()
		if hi < 0 {
			hi += len
		}
	} else {
		hi = max
	}
	// clamp lo to [min;max]
	if lo < min {
		lo = min
	} else if lo > max {
		lo = max
	}
	// clamp hi to [lo;max]
	if hi < lo {
		hi = lo
	} else if hi > max {
		hi = max
	}
	return lo, hi, st
}

// UnsafeSubscript tries to subscript val no matter if it's a list or a map
func UnsafeSubscript(val interface{}, idx interface{}) interface{} {
	switch rval := reflect.ValueOf(val); rval.Kind() {
	case reflect.Slice:
		idx := idx.(int)
		if idx < 0 {
			idx += rval.Len()
		}
		rres := rval.Index(idx)
		if !rres.IsValid() {
			panic(newKindedExtractorError("subscript", fmt.Sprintf("Index out of bounds: %v", idx)))
		}
		return rres.Interface()
	case reflect.Map:
		rres := rval.MapIndex(reflect.ValueOf(idx))
		if !rres.IsValid() {
			panic(newKindedExtractorError("subscript", fmt.Sprintf("Key error: %v", idx)))
		}
		return rres.Interface()
	default:
		if IsTuple(rval) {
			return GetTupleField(rval, idx.(int))
		}
		panic(newKindedExtractorError("subscript", fmt.Sprintf("Unable to subscript %T", val)))
	}
}

// UnsafeSlice tries to slice val no matter if it's a list or a string
func UnsafeSlice(val interface{}, low, high, step OptInt) interface{} {
	switch rval := reflect.ValueOf(val); rval.Kind() {
	case reflect.String:
		return StrSlice(rval.String(), low, high, step)
	case reflect.Slice:
		result := []interface{}{}
		lo, hi, st := InitSliceArgs(rval.Len(), low, high, step)
		if st > 0 {
			for i := lo; i < hi; i += st {
				result = append(result, rval.Index(i).Interface())
			}
		} else {
			for i := hi; i > lo; i += st {
				result = append(result, rval.Index(i).Interface())
			}
		}
		return result
	default:
		panic(newExtractorError(fmt.Sprintf("Unable to slice %T", val)))
	}
}

// UnsafeBinOp tries to apply binary op on left and right according to left's type.
// This method assumes that left and right are unboxed.
func UnsafeBinOp(op string, left interface{}, right interface{}) interface{} {
	if op == "+" {
		switch vleft := left.(type) {
		case string:
			return vleft + right.(string)
		case int:
			return vleft + right.(int)
		case float64:
			return vleft + right.(float64)
		}
	}
	panic(newExtractorError(fmt.Sprintf("Unsupported operation: %T %s %T ", left, op, right)))
}

// UnsafeLen implements Python's len() by type-switching on val
func UnsafeLen(val interface{}) int {
	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.String: // handle Str
		return len(rval.String())
	case reflect.Slice: // handle List, Bytes
		return rval.Len()
	case reflect.Map: // handle SDict, Set
		return rval.Len()
	default:
		if IsTuple(rval) { // handle Tuple
			return rval.NumField()
		} else if ostr, ok := val.(OptString); ok { // handle OptStr
			return len(ostr.Get())
		}
		panic(newExtractorError(fmt.Sprintf("Invalid type passed to len: %T", val)))
	}
}

// DictGet implements python/dict.get
func DictGet(dict SDict, key string, def interface{}) interface{} {
	if val, found := dict[key]; found {
		return val
	}
	return def
}

// DictContains checks if a key is present in a dictionary
func DictContains(dict SDict, key string) bool {
	_, found := dict[key]
	return found
}

// DictUpdate implements python/dict.update
func DictUpdate(dict SDict, other SDict) {
	for k, v := range other {
		dict[k] = v
	}
}

// DictKeys implements python/dict.keys
func DictKeys(dict SDict) []string {
	keys := []string{}
	for key := range dict {
		keys = append(keys, key)
	}
	return keys
}

// DictValues implements python/dict.values
func DictValues(dict SDict) []interface{} {
	values := []interface{}{}
	for _, value := range dict {
		values = append(values, value)
	}
	return values
}

// DictItems implements python/dict.items
func DictItems(dict SDict) []struct {
	Φ0 string
	Φ1 interface{}
} {
	items := []struct {
		Φ0 string
		Φ1 interface{}
	}{}
	for key, value := range dict {
		items = append(items, struct {
			Φ0 string
			Φ1 interface{}
		}{key, value})
	}
	return items
}

// IsTruthy tells whether the supplied interface value holds a truthy value
// (non-empty string/list, True boolean, >0 int, etc.)
func IsTruthy(val interface{}) bool {
	// check sequence types
	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.Slice: // handle List, Bytes
		return rval.Len() > 0
	case reflect.Map: // handle SDict, Set
		return rval.Len() > 0
	}
	// check the rest of the types
	switch v := val.(type) {
	case OptInt:
		return v.IsSet() && IsTruthy(v.Get())
	case OptString:
		return v.IsSet() && IsTruthy(v.Get())
	case OptFloat:
		return v.IsSet() && IsTruthy(v.Get())
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

// IsNone tells whether the supplied interface value holds a Python None
// (empty optional, nil, etc.)
func IsNone(val interface{}) bool {
	switch v := val.(type) {
	case OptInt:
		return !v.IsSet()
	case OptString:
		return !v.IsSet()
	case OptFloat:
		return !v.IsSet()
	default:
		return v == nil
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

// Min returns the minimum of its parameters
func Min(is ...int) int {
	if len(is) == 0 {
		panic(newExtractorError("Min called with no parameters"))
	}
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}
