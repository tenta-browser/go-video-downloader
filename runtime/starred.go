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
 * starred.go: Representation of starred objects and utility for unpacking them
 */

package runtime

import "fmt"

// Starred ..
type Starred interface {
	Object
	Value() Object
}

type starredStruct struct {
	Object
	value Object
}

func (ss *starredStruct) Value() Object {
	return ss.value
}

var _ Starred = (*starredStruct)(nil)

// AsStarred ..
func AsStarred(value Object) Starred {
	return &starredStruct{newObject(ObjectType), value}
}

// Unpack ..
func Unpack(elems ...Object) []Object {
	unpacked := make([]Object, 0, len(elems))
	for _, elem := range elems {
		if s, ok := elem.(Starred); ok {
			seqApply(s.Value(), func(selems []Object) {
				unpacked = append(unpacked, selems...)
			})
		} else {
			unpacked = append(unpacked, elem)
		}
	}
	return unpacked
}

// UnpackIterable ..
func UnpackIterable(o Object, cnt int) Object {
	it := Iter(o)
	unpacked := make([]Object, cnt)
	for i := 0; i < cnt; i++ {
		elem := NextDefault(it, AfterLast)
		if elem == AfterLast {
			panic(RaiseType(ValueErrorType,
				fmt.Sprintf("not enough values to unpack (expected %d, got %d)", cnt, i)))
		}
		unpacked[i] = elem
	}
	if NextDefault(it, AfterLast) != AfterLast {
		panic(RaiseType(ValueErrorType,
			fmt.Sprintf("ValueError: too many values to unpack (expected %d)", cnt)))
	}
	return NewList(unpacked...)
}
