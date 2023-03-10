// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/3/10 9:59
// @Author      : Createitv
// @FileName    : condition_test.go.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	assert.Equal(t, And(1, 1), true)
	assert.Equal(t, And(1, 0), false)
	assert.Equal(t, And(0, 1), false)
	assert.Equal(t, And("", 2), false)
}

func TestBool(t *testing.T) {
	assert.Equal(t, Bool(true), true)
	assert.Equal(t, Bool(false), false)
	assert.Equal(t, Bool(0), false)
	assert.Equal(t, Bool(""), false)
	assert.Equal(t, Bool([]string{}), false)
	assert.Equal(t, Bool(map[string]bool{}), false)
	assert.Equal(t, Bool(map[string]bool{"a": false}), true)
	assert.Equal(t, Bool([]string{"z", "c"}), true)
	assert.Equal(t, Bool([]rune{}), false)
	assert.Equal(t, Bool([3]rune{}), true)

	var ch chan int
	assert.Equal(t, Bool(ch), false)

	var err error
	assert.Equal(t, Bool(err), false)

	assert.Equal(t, Bool(struct {
	}{},
	), false,
	)
}

func TestNand(t *testing.T) {
	assert.Equal(t, Nand(1, 1), false)
	assert.Equal(t, Nand(1, 0), true)
	assert.Equal(t, Nand(0, 1), true)
	assert.Equal(t, Nand(0, 0), true)
}

func TestNor(t *testing.T) {
	assert.Equal(t, Nor(1, 1), false)
	assert.Equal(t, Nor(1, 0), false)
	assert.Equal(t, Nor(0, 1), false)
	assert.Equal(t, Nor(0, 0), true)
}

func TestOr(t *testing.T) {
	assert.Equal(t, Or(1, 1), true)
	assert.Equal(t, Or(0, 1), true)
	assert.Equal(t, Or(1, 0), true)
	assert.Equal(t, Or(0, 0), false)
}

func TestTernary(t *testing.T) {
	assert.Equal(t, Ternary(true, 0, 0), 0)
	assert.Equal(t, Ternary(false, 0, 1), 1)
	assert.Equal(t, Ternary(Bool(""), 0, 3), 3)
}

func TestXnor(t *testing.T) {
	assert.Equal(t, Xnor(1, 1), true)
	assert.Equal(t, Xnor(1, 0), false)
	assert.Equal(t, Xnor(0, 1), false)
	assert.Equal(t, Xnor(0, 0), true)
}

func TestXor(t *testing.T) {
	assert.Equal(t, Xor(1, 1), false)
	assert.Equal(t, Xor(1, 0), true)
	assert.Equal(t, Xor(0, 1), true)
	assert.Equal(t, Xor(0, 0), false)
}
