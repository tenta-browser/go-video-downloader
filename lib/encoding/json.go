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
 * json.go: JSON marshalling and unmarshalling
 */

package encoding

import (
	"encoding/json"
	"fmt"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// JSONLoads implements python/json.loads
var JSONLoads = rnt.NewSimpleFunction("loads", []string{"s"},
	func(args []rnt.Object) rnt.Object {
		s := args[0].(rnt.Str).Value()
		var v interface{}
		if err := json.Unmarshal([]byte(s), &v); err != nil {
			panic(rnt.RaiseType(rnt.ValueErrorType, err.Error()))
		}
		return jsonWrap(v)
	})

// JSONDumps implements python/json.dumps
var JSONDumps = rnt.NewSimpleFunction("dumps", []string{"obj"},
	func(args []rnt.Object) rnt.Object {
		obj := args[0]
		b, err := json.Marshal(jsonUnwrap(obj))
		if err != nil {
			panic(rnt.RaiseType(rnt.TypeErrorType, err.Error()))
		}
		return rnt.NewStr(string(b))
	})

func jsonWrap(v interface{}) rnt.Object {
	if v == nil {
		return rnt.None
	}
	switch vv := v.(type) {
	case string:
		return rnt.NewStr(vv)
	case float64:
		if iv := int(vv); float64(iv) == vv {
			return rnt.NewInt(iv)
		}
		return rnt.NewFloat(vv)
	case bool:
		return rnt.NewBool(vv)
	case []interface{}:
		len := len(vv)
		welems := make([]rnt.Object, len)
		for i, elem := range vv {
			welems[i] = jsonWrap(elem)
		}
		return rnt.NewList(welems...)
	case map[string]interface{}:
		dict := rnt.NewDict()
		for key, val := range vv {
			dict.SetItem(rnt.NewStr(key), jsonWrap(val))
		}
		return dict
	default:
		panic(rnt.RaiseType(rnt.NotImplementedErrorType,
			fmt.Sprintf("cannot wrap type %T", vv)))
	}
}

func jsonUnwrap(o rnt.Object) interface{} {
	if o == rnt.None {
		return nil
	}
	if o.IsInstance(rnt.StrType) {
		return o.(rnt.Str).Value()
	}
	if o.IsInstance(rnt.IntType) {
		return o.(rnt.Int).Value()
	}
	if o.IsInstance(rnt.FloatType) {
		return o.(rnt.Float).Value()
	}
	if o.IsInstance(rnt.BoolType) {
		return o.(rnt.Int).Value() != 0
	}
	if o.IsInstance(rnt.ListType) {
		return jsonUnwrapSeq(o.(rnt.List).Elems())
	}
	if o.IsInstance(rnt.TupleType) {
		return jsonUnwrapSeq(o.(rnt.Tuple).Elems())
	}
	if o.IsInstance(rnt.DictType) {
		m := make(map[string]interface{})
		it := o.(rnt.Dict).Iterator()
		for e := it(); e != nil; e = it() {
			key := e.Key
			value := e.Value
			var skey string
			if key == rnt.None {
				skey = "null"
			} else if key == rnt.True {
				skey = "true"
			} else if key == rnt.False {
				skey = "false"
			} else if !key.IsInstance(rnt.StrType) &&
				!key.IsInstance(rnt.IntType) &&
				!key.IsInstance(rnt.FloatType) {
				panic(rnt.RaiseType(rnt.TypeErrorType,
					"keys must be str, int, float, bool or None, not object"))
			} else {
				skey = rnt.ToStr(key).Value()
			}
			m[skey] = jsonUnwrap(value)
		}
		return m
	}
	panic(rnt.RaiseType(rnt.TypeErrorType,
		fmt.Sprintf("cannot unwrap type %s", o.Type().Name())))
}

func jsonUnwrapSeq(os []rnt.Object) []interface{} {
	len := len(os)
	res := make([]interface{}, len)
	for i, o := range os {
		res[i] = jsonUnwrap(o)
	}
	return res
}
