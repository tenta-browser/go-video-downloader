package runtime

import "fmt"

// CastToInt casts to int, handles boxing, panics with extractorError
func CastToInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return v
	case OptInt:
		return v.Get()
	default:
		panic(NewCastError(val, "int"))
	}
}

// CastToString casts to string, handles boxing, panics with extractorError
func CastToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case OptString:
		return v.Get()
	default:
		panic(NewCastError(val, "str"))
	}
}

// CastToOptInt casts to OptInt, handles boxing, panics with extractorError
func CastToOptInt(val interface{}) OptInt {
	switch v := val.(type) {
	case nil:
		return OptInt{}
	case int:
		return AsOptInt(v)
	case OptInt:
		return v
	default:
		panic(NewCastError(val, "OptInt"))
	}
}

// CastToOptString casts to OptString, handles boxing, panics with extractorError
func CastToOptString(val interface{}) OptString {
	switch v := val.(type) {
	case nil:
		return OptString{}
	case string:
		return AsOptString(v)
	case OptString:
		return v
	default:
		panic(NewCastError(val, "OptString"))
	}
}

// NewCastError creates extractorError's for failed casts
func NewCastError(val interface{}, destType string) error {
	return newExtractorError(fmt.Sprintf("Cannot cast %T to %s", val, destType))
}
