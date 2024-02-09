package mirror

import "reflect"

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
