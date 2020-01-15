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
 * literal.go: Memory efficient implementation of string, integer and float types,
 *             to represent literal values.
 */

package runtime

import (
	"fmt"
	"reflect"
)

type (
	// StrLiteral ..
	StrLiteral string
	// IntLiteral ..
	IntLiteral int
	// FloatLiteral ..
	FloatLiteral float64
)

var (
	strLiteralImpl   = newObject(StrType)
	intLiteralImpl   = newObject(IntType)
	floatLiteralImpl = newObject(FloatType)
)

// ---- String literal ----

// Type ..
func (sc StrLiteral) Type() Type {
	return strLiteralImpl.Type()
}

// Dict ..
func (sc StrLiteral) Dict() Dict {
	return strLiteralImpl.Dict()
}

func (sc StrLiteral) setDict(d Dict) {
	strLiteralImpl.setDict(d)
}

// IsInstance ..
func (sc StrLiteral) IsInstance(t Type) bool {
	return strLiteralImpl.IsInstance(t)
}

// Value ..
func (sc StrLiteral) Value() string {
	return string(sc)
}

// ---- Int literal ----

// Type ..
func (ic IntLiteral) Type() Type {
	return intLiteralImpl.Type()
}

// Dict ..
func (ic IntLiteral) Dict() Dict {
	return intLiteralImpl.Dict()
}

func (ic IntLiteral) setDict(d Dict) {
	intLiteralImpl.setDict(d)
}

// IsInstance ..
func (ic IntLiteral) IsInstance(t Type) bool {
	return intLiteralImpl.IsInstance(t)
}

// Value ..
func (ic IntLiteral) Value() int {
	return int(ic)
}

// ---- Float literal ----

// Type ..
func (ic FloatLiteral) Type() Type {
	return floatLiteralImpl.Type()
}

// Dict ..
func (ic FloatLiteral) Dict() Dict {
	return floatLiteralImpl.Dict()
}

func (ic FloatLiteral) setDict(d Dict) {
	floatLiteralImpl.setDict(d)
}

// IsInstance ..
func (ic FloatLiteral) IsInstance(t Type) bool {
	return floatLiteralImpl.IsInstance(t)
}

// Value ..
func (ic FloatLiteral) Value() float64 {
	return float64(ic)
}

// ---- Dict literal ----

// DictLiteral ..
func DictLiteral(m interface{}) Dict {
	type converter func(reflect.Value) Object

	var oi Object
	ot := reflect.TypeOf(&oi).Elem()

	getConverter := func(t reflect.Type) converter {
		if t == ot {
			return func(v reflect.Value) Object {
				return v.Interface().(Object)
			}
		}
		switch t.Kind() {
		case reflect.String:
			return func(v reflect.Value) Object {
				return NewStr(v.String())
			}
		case reflect.Int:
			return func(v reflect.Value) Object {
				return NewInt(int(v.Int()))
			}
		case reflect.Float64:
			return func(v reflect.Value) Object {
				return NewFloat(v.Float())
			}
		default:
			panic(fmt.Sprintf("Unhandled type: %v", t))
		}
	}

	v := reflect.ValueOf(m)
	t := v.Type()

	keyConv := getConverter(t.Key())
	valConv := getConverter(t.Elem())

	d := newDict(DictType)
	it := v.MapRange()
	for it.Next() {
		d.SetItem(keyConv(it.Key()), valConv(it.Value()))
	}

	return d
}
