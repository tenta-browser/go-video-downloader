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
 * bytes.go: Implementation of Python's bytes type
 */

package runtime

import (
	"bytes"
	"fmt"
)

// Bytes ..
type Bytes interface {
	Object
	Value() []byte
}

type bytesStruct struct {
	Object
	value []byte
}

func (bs *bytesStruct) Value() []byte {
	return bs.value
}

var _ Bytes = (*bytesStruct)(nil)

// NewBytes ..
func NewBytes(value ...byte) Bytes {
	return &bytesStruct{newObject(BytesType), value}
}

// BytesType ..
var BytesType = newBuiltinType("bytes")

func bytesNew(t Type, args Args, kwArgs KWArgs) Object {
	// bytes([source[, encoding[, errors]]])
	checkFunctionArgsMin("bytes", args, kwArgs, 0, ObjectType, StrType, StrType)
	argc := len(args)
	if argc == 0 {
		// empty
		return NewBytes()
	}
	if args[0].IsInstance(IntType) {
		// zero-filled with given length
		len := args[0].(Int).Value()
		return NewBytes(make([]byte, len)...)
	}
	if args[0].IsInstance(StrType) {
		if argc < 2 {
			panic(RaiseType(TypeErrorType, "string argument without an encoding"))
		}
		encoding := args[1].(Str).Value()
		errors := NewStr(codingErrorsStrict)
		if argc >= 3 {
			errors = args[2].(Str)
		}
		bytes, _ := GetCodec(encoding).Encode(args[0], errors)
		return bytes
	}
	// it must be an iterable
	value := []byte{}
	seqForEach(args[0], func(item Object) {
		if !item.IsInstance(IntType) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"'%s' object cannot be interpreted as an integer", item.Type().Name())))
		}
		i := item.(Int).Value()
		if i < 0 || i >= 256 {
			panic(RaiseType(ValueErrorType, "bytes must be in range(0, 256)"))
		}
		value = append(value, byte(i))
	})
	return NewBytes(value...)
}

func bytesHash(o Object) Object {
	b := o.(Bytes)
	h := newHash()
	h.Write(b.Value())
	return NewInt(sumHash(h))
}

func bytesRepr(o Object) Object {
	b := o.(Bytes)
	return NewStr("b" + Repr(NewStr(string(b.Value()))).Value())
}

// copy of seqGetItem, using bytes instead of Objects
func bytesGetItem(o Object, index Object) Object {
	t := o.Type()
	value := o.(Bytes).Value()
	if index.IsInstance(IntType) {
		i := seqCheckedIndex(t, len(value), index.(Int).Value())
		return NewInt(int(value[i]))
	} else if index.IsInstance(SliceType) {
		start, stop, step, len := index.(Slice).CalcSlice(len(value))
		if step == 1 {
			return NewBytes(value[start:stop]...)
		}
		sliced := make([]byte, len)
		i := 0
		for j := start; j != stop; j += step {
			sliced[i] = value[j]
			i++
		}
		return NewBytes(sliced...)
	} else {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"%s indices must be integers or slices, not %s",
			t.Name(), index.Type().Name())))
	}
}

func bytesLen(o Object) Object {
	b := o.(Bytes)
	return NewInt(len(b.Value()))
}

func bytesContains(o Object, needle Object) Object {
	b := o.(Bytes).Value()
	if needle.IsInstance(IntType) {
		i := needle.(Int).Value()
		if i < 0 || i >= 256 {
			panic(RaiseType(ValueErrorType, "byte must be in range(0, 256)"))
		}
		n := byte(i)
		for _, bb := range b {
			if bb == n {
				return True
			}
		}
		return False
	}
	if needle.IsInstance(BytesType) {
		nb := needle.(Bytes).Value()
		return NewBool(bytes.Contains(b, nb))
	}
	panic(RaiseType(TypeErrorType, fmt.Sprintf(
		"a bytes-like object is required, not '%s'", needle.Type().Name())))
}

func bytesCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(BytesType) {
		return NotImplemented
	}
	cmp := bytes.Compare(a.(Bytes).Value(), b.(Bytes).Value())
	var r bool
	switch op {
	case compareEq:
		r = cmp == 0
	case compareNe:
		r = cmp != 0
	case compareLt:
		r = cmp < 0
	case compareGt:
		r = cmp > 0
	case compareLe:
		r = cmp <= 0
	case compareGe:
		r = cmp >= 0
	}
	return NewBool(r)
}

func bytesDecode(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("decode", args, kwArgs, 1, BytesType, StrType, StrType)
	argc := len(args)
	encoding := "utf-8"
	if argc >= 2 {
		encoding = args[1].(Str).Value()
	}
	errors := NewStr(codingErrorsStrict)
	if argc >= 3 {
		errors = args[2].(Str)
	}
	s, _ := GetCodec(encoding).Decode(args[0], errors)
	return s
}

func bytesStartswith(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("startswith", args, kwArgs, BytesType, BytesType)
	b := args[0].(Bytes).Value()
	prefix := args[1].(Bytes).Value()
	return NewBool(bytes.HasPrefix(b, prefix))
}

func initBytesType(slots *typeSlots, dict map[string]Object) {
	dict["decode"] = newBuiltinFunction("decode", bytesDecode)
	dict["startswith"] = newBuiltinFunction("startswith", bytesStartswith)
	slots.New = bytesNew
	slots.Hash = bytesHash
	slots.Repr = bytesRepr
	slots.GetItem = bytesGetItem
	slots.Len = bytesLen
	slots.Contains = bytesContains
	slots.Eq = makeCompareFn(bytesCompare, compareEq)
	slots.Ne = makeCompareFn(bytesCompare, compareNe)
	slots.Lt = makeCompareFn(bytesCompare, compareLt)
	slots.Gt = makeCompareFn(bytesCompare, compareGt)
	slots.Le = makeCompareFn(bytesCompare, compareLe)
	slots.Ge = makeCompareFn(bytesCompare, compareGe)
}
