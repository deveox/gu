package mirror

import "reflect"

func DerefType(t reflect.Type) reflect.Type {
	for {
		switch t.Kind() {
		case reflect.Ptr, reflect.Interface:
			t = t.Elem()
		default:
			return t
		}
	}
}
