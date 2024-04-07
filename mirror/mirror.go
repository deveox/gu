package mirror

import (
	"reflect"
)

func DerefType(t reflect.Type) reflect.Type {
	for {
		switch t.Kind() {
		case reflect.Ptr:
			t = t.Elem()
		default:
			return t
		}
	}
}

func DerefValue(v reflect.Value) reflect.Value {
	for {
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface:
			v = v.Elem()
		default:
			return v
		}
	}
}

// New creates a new instance of the given generic type.
//
// v := New[int]() // v is int
// v := New[*int]() // v is *int
// v := New[**int]() is not supported
func New[T any]() (T, reflect.Type) {
	t := reflect.TypeFor[T]()
	i := reflect.New(DerefType(t))
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		return i.Interface().(T), t
	default:
		return i.Elem().Interface().(T), t
	}
}
