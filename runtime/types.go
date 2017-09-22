package runtime

// OptString represents an optional string
type OptString struct {
	string
	set bool
}

// AsOptString converts a string to an optional string
func AsOptString(value string) OptString {
	return OptString{value, true}
}

// Get returns the value of the optional object
func (obj OptString) Get() string {
	return obj.string
}

// GetOrDef returns the value of the optional object is it's available
// or the default value otherwise
func (obj OptString) GetOrDef(def string) string {
	if !obj.IsSet() {
		return def
	}
	return obj.string
}

// Set sets the value of the optional object
func (obj *OptString) Set(value string) {
	obj.string = value
	obj.set = true
}

// IsSet returns whether the optional object has a value
func (obj OptString) IsSet() bool {
	return obj.set
}

// OptInt represents an optional int
type OptInt struct {
	int
	set bool
}

// AsOptInt converts an int to an optional int
func AsOptInt(value int) OptInt {
	return OptInt{value, true}
}

// Get returns the value of the optional object
func (obj OptInt) Get() int {
	return obj.int
}

// GetOrDef returns the value of the optional object is it's available
// or the default value otherwise
func (obj OptInt) GetOrDef(def int) int {
	if !obj.IsSet() {
		return def
	}
	return obj.int
}

// Set sets the value of the optional object
func (obj *OptInt) Set(value int) {
	obj.int = value
	obj.set = true
}

// IsSet returns whether the optional object has a value
func (obj OptInt) IsSet() bool {
	return obj.set
}

type noDefaultType struct{}

// NoDefault is the simulation of youtube-dl's hack NO_DEFAULT = object() (utils.py),
// which is used in some places as a marker value (where None meant something else)
var NoDefault = noDefaultType{}

func isNoDefault(something interface{}) bool {
	_, ok := something.(noDefaultType)
	return ok
}
