package mirror

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDerefType(t *testing.T) {
	typ := reflect.TypeOf(&struct{}{})
	require.Equal(t, reflect.TypeOf(struct{}{}), DerefType(typ))
	v := 0
	v1 := &v
	v2 := &v1
	typ = reflect.TypeOf(v2)
	require.Equal(t, reflect.TypeOf(0), DerefType(typ))
	var iv interface{} = 0
	typ = reflect.TypeOf(&iv)
	require.Equal(t, reflect.TypeOf(0), DerefType(typ))
}
