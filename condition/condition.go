// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/3/10 9:23
// @Author      : Createitv
// @FileName    : condition.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

// Package condition contains some functions for conditional judgment. eg. And, Or, TernaryOperator ...
// The implementation of this package refers to the implementation of carlmjohnson's truthy package, you may find more
package condition

import "reflect"

// Ternary operator is a conditional operator that likes three operands: a condition followed by a question mark (?),
// Ternary returns ifVal if cond is true, elseVal otherwise.
func Ternary[T, U any](cond U, ifVal, elseVal T) T {
	if Bool(cond) {
		return ifVal
	}
	return elseVal
}

// Bool returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
//
// Note that the usual Go type system caveats apply around a nil pointer value not being a nil interface value.
//
// In benchmark testing,
// ValueAny is approximately 25 times slower than Value,
// and 50 times slower than native Go comparisons.
func Bool[T any](v T) bool {
	switch m := any(v).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&v)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.String:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}

// And returns true if both a and b are truthy.
func And[T, U any](left T, right U) bool {
	return Bool(left) && Bool(right)
}

// Or returns true if either a or b are truthy.
func Or[T, U any](left T, right U) bool {
	return Bool(left) || Bool(right)
}

// Xor returns true if a or b but not both is truthy.
func Xor[T, U any](left T, right U) bool {
	valA := Bool(left)
	valB := Bool(right)
	return (valA || valB) && valA != valB
}

// Nor returns true if neither a nor b is truthy.
func Nor[T, U any](left T, right U) bool {
	return !(Bool(left) || Bool(right))
}

// Xnor returns true if both a and b or neither a nor b are truthy.
func Xnor[T, U any](left T, right U) bool {
	valA := Bool(left)
	valB := Bool(right)
	return (valA && valB) || (!valA && !valB)
}

// Nand returns false if both a and b are truthy.
func Nand[T, U any](left T, b U) bool {
	return !Bool(left) || !Bool(b)
}
