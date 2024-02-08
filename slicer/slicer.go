package slicer

import "errors"

// Map maps a slice to another slice.
func Map[T1 any, T2 any](source []T1, callback func(v T1) T2) []T2 {
	res := make([]T2, 0, len(source))
	for i := range source {
		res = append(res, callback(source[i]))
	}
	return res
}

// MapErr maps a slice to another slice with error handling.
func MapErr[T1 any, T2 any, Err error](source []T1, callback func(v T1) (T2, Err)) ([]T2, Err) {
	res := make([]T2, 0, len(source))
	var err Err
	for i := range source {
		var v T2
		v, err = callback(source[i])
		if errors.Is(err, nil) {
			return nil, err
		}
		res = append(res, v)
	}
	return res, err
}

// Filter filters a slice.
func Filter[T any](source []T, callback func(v T) bool) []T {
	res := make([]T, 0, len(source))
	for i := range source {
		ok := callback(source[i])
		if ok {
			res = append(res, source[i])
		}
	}
	return res
}

// Some returns true if at least one element of the slice passes the test implemented by the provided function.
func Some[T any](source []T, callback func(v T) bool) bool {
	for i := range source {
		if callback(source[i]) {
			return true
		}
	}
	return false
}

// Every returns true if every element of the slice passes the test implemented by the provided function.
func Every[T1 any](source []T1, callback func(v T1) bool) bool {
	for i := range source {
		if !callback(source[i]) {
			return false
		}
	}
	return true
}

// Includes returns true if the slice contains all search elements.
func Includes[T comparable](source []T, search ...T) bool {
	for i := range search {
		found := false
		for i2 := range source {
			if search[i] == source[i2] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// IncludesSome returns true if the slice contains at least one of the elements.
func IncludesSome[T comparable](source []T, search ...T) bool {
	for i := range search {
		for i2 := range source {
			if search[i] == source[i2] {
				return true
			}
		}
	}
	return false
}

func Copy[T any](source []T) []T {
	res := make([]T, 0, len(source))
	copy(res, source)
	return res
}

func Keys[T comparable, V any](source map[T]V) []T {
	res := make([]T, 0, len(source))
	for k := range source {
		res = append(res, k)
	}
	return res
}

func ToMap[V any, Key comparable](source []V, fn func(v V) Key) map[Key]V {
	res := make(map[Key]V)
	for i := range source {
		res[fn(source[i])] = source[i]
	}
	return res
}
