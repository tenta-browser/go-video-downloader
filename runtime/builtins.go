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
 * builtins.go: Implementation of Python's builtins
 */

package runtime

import "fmt"

// BuiltinPrint ..
var BuiltinPrint = newBuiltinFunction("print", func(args Args, kwArgs KWArgs) Object {
	checkFunctionKWArgs("print", kwArgs, map[string]Type{
		"sep": StrType,
		"end": StrType,
	})
	sep := kwArgs.get("sep", NewStr(" ")).(Str)
	end := kwArgs.get("end", NewStr("\n")).(Str)
	Print(args, sep, end)
	return None
})

// BuiltinHash ..
var BuiltinHash = newBuiltinFunction("hash", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("hash", args, kwArgs, ObjectType)
	return Hash(args[0])
})

// BuiltinRepr ..
var BuiltinRepr = newBuiltinFunction("repr", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("repr", args, kwArgs, ObjectType)
	return Repr(args[0])
})

// BuiltinLen ..
var BuiltinLen = newBuiltinFunction("len", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("len", args, kwArgs, ObjectType)
	return Len(args[0])
})

// BuiltinIter ..
var BuiltinIter = newBuiltinFunction("iter", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("iter", args, kwArgs, 1, ObjectType, ObjectType)
	if len(args) == 2 {
		// TODO iter(callable, sentinel)
		panic(RaiseType(NotImplementedErrorType, "iter(callable, sentinel) not implemented"))
	}
	return Iter(args[0])
})

// BuiltinNext ..
var BuiltinNext = newBuiltinFunction("next", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("next", args, kwArgs, 1, ObjectType, ObjectType)
	if len(args) == 2 {
		return NextDefault(args[0], args[1])
	}
	return Next(args[0])
})

// BuiltinIsInstance ..
var BuiltinIsInstance = newBuiltinFunction("isinstance", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("isinstance", args, kwArgs, ObjectType, ObjectType)
	return NewBool(IsInstance(args[0], args[1]))
})

// BuiltinIsSubclass ..
var BuiltinIsSubclass = newBuiltinFunction("issubclass", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("issubclass", args, kwArgs, ObjectType, ObjectType)
	return NewBool(IsSubclass(args[0], args[1]))
})

// BuiltinAll ..
var BuiltinAll = newBuiltinFunction("all", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("all", args, kwArgs, ObjectType)
	hasFalse := seqFindFirst(args[0], func(o Object) bool {
		return !IsTrue(o)
	})
	return NewBool(!hasFalse)
})

// BuiltinAny ..
var BuiltinAny = newBuiltinFunction("any", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("any", args, kwArgs, ObjectType)
	hasTrue := seqFindFirst(args[0], func(o Object) bool {
		return IsTrue(o)
	})
	return NewBool(hasTrue)
})

// BuiltinCallable ..
var BuiltinCallable = newBuiltinFunction("callable", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("callable", args, kwArgs, ObjectType)
	callFn := args[0].Type().Slots().Call
	return NewBool(callFn != nil)
})

// BuiltinSorted ..
var BuiltinSorted = newBuiltinFunction("sorted", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("sorted", args, nil, ObjectType)
	checkFunctionKWArgs("sorted", kwArgs, map[string]Type{
		"key":     ObjectType,
		"reverse": IntType,
	})
	l := Cal(ListType, args[0])
	listSort(Args{l}, kwArgs)
	return l
})

// BuiltinGetAttr ..
var BuiltinGetAttr = newBuiltinFunction("getattr", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("getattr", args, kwArgs, 2, ObjectType, StrType, ObjectType)
	o := args[0]
	name := args[1].(Str).Value()
	var def Object
	if len(args) >= 3 {
		def = args[2]
	}
	return GetAttr(o, name, def)
})

// BuiltinSetAttr ..
var BuiltinSetAttr = newBuiltinFunction("setattr", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("setattr", args, kwArgs, ObjectType, StrType, ObjectType)
	o := args[0]
	name := args[1].(Str).Value()
	value := args[2]
	SetAttr(o, name, value)
	return None
})

// BuiltinHasAttr ..
var BuiltinHasAttr = newBuiltinFunction("hasattr", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("hasattr", args, kwArgs, ObjectType, StrType)
	o := args[0]
	name := args[1].(Str).Value()
	return NewBool(HasAttr(o, name))
})

// BuiltinOpen ..
var BuiltinOpen = newBuiltinFunction("open", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("open", args, nil, ObjectType)
	panic(RaiseType(NotImplementedErrorType, "open() not implemented"))
})

// BuiltinMin ..
var BuiltinMin = newBuiltinFunction("min", func(args Args, kwArgs KWArgs) Object {
	return builtinMinMax(false, args, kwArgs)
})

// BuiltinMax ..
var BuiltinMax = newBuiltinFunction("max", func(args Args, kwArgs KWArgs) Object {
	return builtinMinMax(true, args, kwArgs)
})

func builtinMinMax(isMax bool, args Args, kwArgs KWArgs) Object {
	var fname string
	if isMax {
		fname = "max"
	} else {
		fname = "min"
	}
	checkFunctionVarArgs(fname, args, nil, ObjectType)
	checkFunctionKWArgs(fname, kwArgs, map[string]Type{
		"key":     ObjectType,
		"default": ObjectType,
	})
	keyFn := kwArgs.get("key", nil)
	def := kwArgs.get("default", nil)
	var resKey, resVal Object
	applyItem := func(itemVal Object) {
		itemKey := itemVal
		if keyFn != nil {
			itemKey = Cal(keyFn, itemVal)
		}
		if resVal == nil ||
			(isMax && IsTrue(Gt(itemKey, resKey))) ||
			(!isMax && IsTrue(Lt(itemKey, resKey))) {
			resKey = itemKey
			resVal = itemVal
		}
	}
	if len(args) == 1 {
		if def != nil {
			applyItem(def)
		}
		seqForEach(args[0], applyItem)
		if resVal == nil {
			panic(RaiseType(ValueErrorType, fmt.Sprintf(
				"%s() arg is an empty sequence", fname)))
		}
	} else {
		for _, arg := range args {
			applyItem(arg)
		}
	}
	return resVal
}

// BuiltinChr ..
var BuiltinChr = newBuiltinFunction("chr", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("ord", args, kwArgs, IntType)
	i := args[0].(Int).Value()
	if i < 0 || i >= 0x110000 {
		panic(RaiseType(ValueErrorType, "chr() arg not in range(0x110000)"))
	}
	return NewStr(string(rune(i)))
})

// BuiltinOrd ..
var BuiltinOrd = newBuiltinFunction("ord", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("ord", args, kwArgs, ObjectType)
	o := args[0]
	errLen := "ord() expected a character, but string of length %d found"
	if o.IsInstance(StrType) {
		r := []rune(o.(Str).Value())
		if len(r) != 1 {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(errLen, len(r))))
		}
		return NewInt(int(r[0]))
	} else if o.IsInstance(BytesType) {
		b := o.(Bytes).Value()
		if len(b) != 1 {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(errLen, len(b))))
		}
		return NewInt(int(b[0]))
	} else {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"ord() expected string of length 1, but %s found", o.Type().Name())))
	}
})

// BuiltinSum ..
var BuiltinSum = newBuiltinFunction("sum", func(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("sum", args, kwArgs, 1, ObjectType, ObjectType)
	var res Object
	if len(args) == 1 {
		res = NewInt(0)
	} else {
		res = args[1]
	}
	seqForEach(args[0], func(o Object) {
		res = Add(res, o)
	})
	return res
})

type builtinTypeState int

const (
	builtinTypeStateNotReady builtinTypeState = iota
	builtinTypeStateIniting
	builtinTypeStateReady
)

type builtinTypeInfo struct {
	init  func(*typeSlots, map[string]Object)
	dict  map[string]Object
	state builtinTypeState
}

var builtinTypes = map[Type]*builtinTypeInfo{
	TypeType:                {init: initTypeType},
	ObjectType:              {init: initObjectType},
	FunctionType:            {init: initFunctionType},
	NoneType:                {init: initNoneType},
	NotImplementedType:      {init: initNotImplementedType},
	IntType:                 {init: initIntType},
	FloatType:               {init: initFloatType},
	StrType:                 {init: initStrType},
	TupleType:               {init: initTupleType},
	ListType:                {init: initListType},
	DictType:                {init: initDictType},
	SetType:                 {init: initSetType},
	BoolType:                {init: initBoolType},
	BytesType:               {init: initBytesType},
	seqIteratorType:         {init: initSeqIteratorType},
	ZipIteratorType:         {init: initZipIteratorType},
	MapIteratorType:         {init: initMapIteratorType},
	FilterIteratorType:      {init: initFilterIteratorType},
	EnumerateIteratorType:   {init: initEnumerateIteratorType},
	ReversedIteratorType:    {init: initReversedIteratorType},
	RangeType:               {init: initRangeType},
	RangeIteratorType:       {init: initRangeIteratorType},
	GeneratorType:           {init: initGeneratorType},
	SliceType:               {init: initSliceType},
	PropertyType:            {init: initPropertyType},
	MethodType:              {init: initMethodType},
	StaticMethodType:        {init: initStaticMethodType},
	ClassMethodType:         {init: initClassMethodType},
	SuperType:               {init: initSuperType},
	NativeType:              {init: initNativeType},
	BaseExceptionType:       {init: initBaseExceptionType},
	ExceptionType:           {},
	TypeErrorType:           {},
	ValueErrorType:          {},
	RuntimeErrorType:        {},
	NotImplementedErrorType: {},
	LookupErrorType:         {},
	KeyErrorType:            {},
	IndexErrorType:          {},
	StopIterationType:       {},
	AttributeErrorType:      {},
	AssertionErrorType:      {},
	UnicodeErrorType:        {},
	UnicodeEncodeErrorType:  {},
	UnicodeDecodeErrorType:  {},
	OSErrorType:             {},
	ArithmeticErrorType:     {},
	OverflowErrorType:       {},
	ZeroDivisionErrorType:   {},
}

func initBuiltinTypes() {
	for t := range builtinTypes {
		initBuiltinType(t)
	}
	// Note: dicts have to be set in a separate pass,
	// because the dict type is not ready earlier.
	for t, info := range builtinTypes {
		d := NewDict()
		for name, obj := range info.dict {
			d.SetItem(NewStr(name), obj)
		}
		t.setDict(d)
	}
}

func initBuiltinType(t Type) {
	info := builtinTypes[t]
	switch info.state {
	case builtinTypeStateReady:
		return
	case builtinTypeStateIniting:
		panic(fmt.Sprintf("Initialization cycle for %s", t.Name()))
	}
	info.state = builtinTypeStateIniting
	// init bases first
	for _, base := range t.Bases() {
		initBuiltinType(base)
	}
	// init pre-dict and slots
	info.dict = make(map[string]Object)
	if info.init != nil {
		info.init(t.Slots(), info.dict)
	}
	// populate dict from slots
	fillDictFromSlots(t, info.dict)
	// process inheritance
	postInitType(t)
	info.state = builtinTypeStateReady
}

func init() {
	initBuiltinTypes()
}
