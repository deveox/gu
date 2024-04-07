package mirror

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDerefType(t *testing.T) {
	typ := reflect.TypeOf(&struct{}{})
	require.Equal(t, reflect.TypeOf(struct{}{}).Name(), DerefType(typ).Name())
	v := 0
	v1 := &v
	v2 := &v1
	typ = reflect.TypeOf(v2)
	require.Equal(t, reflect.TypeOf(v).Name(), DerefType(typ).Name())
	var iv interface{} = 0
	typ = reflect.TypeOf(&iv)
	require.Equal(t, reflect.TypeFor[interface{}]().Name(), DerefType(typ).Name())
}

func TestNew(t *testing.T) {
	type MyStruct struct{}

	actual, typ := New[MyStruct]()
	assert.Equal(t, reflect.TypeOf(MyStruct{}).Name(), typ.Name(), "New() should return the type of the given type")
	assert.NotNil(t, actual, "New() should not return nil when called with a non-pointer type")
	assert.Equal(t, actual, MyStruct{}, "New() should return a zero value of the given type")
	actual2, typ := New[*MyStruct]()
	assert.Equal(t, reflect.TypeOf(&MyStruct{}).Name(), typ.Name(), "New() should return the type of the given type")
	assert.NotNil(t, actual2, "New() should not return nil when called with a pointer type")
	assert.Equal(t, actual2, &MyStruct{}, "New() should return a zero value of the given type")
}
