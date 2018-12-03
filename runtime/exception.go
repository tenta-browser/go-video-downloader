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
 * exception.go: Implementation of Python's builtin exception types
 */

package runtime

import "fmt"

// BaseException ..
type BaseException interface {
	Object
	Args() Tuple
	setArgs(Tuple)
	String() string
}

type baseExceptionStruct struct {
	Object
	args Tuple
}

func (bes *baseExceptionStruct) Args() Tuple {
	return bes.args
}

func (bes *baseExceptionStruct) setArgs(args Tuple) {
	bes.args = args
}

func (bes *baseExceptionStruct) String() string {
	return fmt.Sprintf("%s: %s", bes.Type().Name(), ToStr(bes).Value())
}

var _ BaseException = (*baseExceptionStruct)(nil)

func newBaseException(t Type) BaseException {
	return &baseExceptionStruct{Object: newObject(t)}
}

// BaseExceptionType ..
var BaseExceptionType = newBuiltinType("BaseException")

func baseExceptionNew(t Type, args Args, kwArgs KWArgs) Object {
	return newBaseException(t)
}

func baseExceptionInit(o Object, args Args, kwArgs KWArgs) Object {
	be := o.(BaseException)
	be.setArgs(NewTuple(args...))
	return None
}

func baseExceptionRepr(o Object) Object {
	be := o.(BaseException)
	args := be.Args()
	var argsRepr string
	if args == nil {
		argsRepr = ""
	} else if len(args.Elems()) != 1 {
		argsRepr = Repr(args).Value()
	} else {
		argsRepr = fmt.Sprintf("(%s)", Repr(args.Elems()[0]).Value())
	}
	return NewStr(be.Type().Name() + argsRepr)
}

func baseExceptionStr(o Object) Object {
	be := o.(BaseException)
	args := be.Args()
	if args == nil || len(args.Elems()) == 0 {
		return NewStr("")
	}
	if len(args.Elems()) == 1 {
		return ToStr(args.Elems()[0])
	}
	return ToStr(args)
}

func initBaseExceptionType(slots *typeSlots, dict map[string]Object) {
	slots.New = baseExceptionNew
	slots.Init = baseExceptionInit
	slots.Repr = baseExceptionRepr
	slots.Str = baseExceptionStr
}

var (
	// ExceptionType ..
	ExceptionType = newBuiltinType("Exception", BaseExceptionType)
	// TypeErrorType ..
	TypeErrorType = newBuiltinType("TypeError", ExceptionType)
	// ValueErrorType ..
	ValueErrorType = newBuiltinType("ValueError", ExceptionType)
	// RuntimeErrorType ..
	RuntimeErrorType = newBuiltinType("RuntimeError", ExceptionType)
	// NotImplementedErrorType ..
	NotImplementedErrorType = newBuiltinType("NotImplementedError", RuntimeErrorType)
	// LookupErrorType ..
	LookupErrorType = newBuiltinType("LookupError", ExceptionType)
	// KeyErrorType ..
	KeyErrorType = newBuiltinType("KeyError", LookupErrorType)
	// IndexErrorType ..
	IndexErrorType = newBuiltinType("IndexError", LookupErrorType)
	// StopIterationType ..
	StopIterationType = newBuiltinType("StopIteration", ExceptionType)
	// AttributeErrorType ..
	AttributeErrorType = newBuiltinType("AttributeError", ExceptionType)
	// AssertionErrorType ..
	AssertionErrorType = newBuiltinType("AssertionError", ExceptionType)
	// ImportErrorType ..
	ImportErrorType = newBuiltinType("ImportError", ExceptionType)
	// UnicodeErrorType ..
	UnicodeErrorType = newBuiltinType("UnicodeError", ValueErrorType)
	// UnicodeEncodeErrorType ..
	UnicodeEncodeErrorType = newBuiltinType("UnicodeEncodeError", UnicodeErrorType)
	// UnicodeDecodeErrorType ..
	UnicodeDecodeErrorType = newBuiltinType("UnicodeDecodeError", UnicodeErrorType)
	// OSErrorType ..
	OSErrorType = newBuiltinType("OSError", ExceptionType)
	// ArithmeticErrorType ..
	ArithmeticErrorType = newBuiltinType("ArithmeticError", ExceptionType)
	// OverflowErrorType ..
	OverflowErrorType = newBuiltinType("OverflowError", ArithmeticErrorType)
	// ZeroDivisionErrorType ..
	ZeroDivisionErrorType = newBuiltinType("ZeroDivisionError", ArithmeticErrorType)
)
