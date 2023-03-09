// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/3/9 22:17
// @Author      : Createitv
// @FileName    : random_test.go.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package random

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIdV4(t *testing.T) {

	uuid, err := UUIdV4()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	isUUiDV4 := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	assert.Equal(t, true, isUUiDV4.MatchString(uuid), "uuid 正则格式需要相匹配")
}

func TestRandBetween(t *testing.T) {
	r1 := RandBetween(1, 10)
	assert.GreaterOrEqual(t, r1, 1)
	assert.Less(t, r1, 10)

	r2 := RandBetween(1, 1)
	assert.Equal(t, 1, r2)

	r3 := RandBetween(10, 1)
	assert.GreaterOrEqual(t, r1, 1)
	assert.Less(t, r3, 10)
}

func TestRandFromString(t *testing.T) {
	c := RandFromString(Alphabet, 10)
	for i := 0; i < len(c); i++ {
		assert.Contains(t, Alphabet, string(c[i]))
	}

	s := "213sadkasj209m231-a."
	z := RandFromString(s, 10)
	for i := 0; i < len(c); i++ {
		assert.Contains(t, s, string(z[i]))
	}
}

func TestRandNumber(t *testing.T) {
	r := RandNumber(10)
	for i := 0; i < len(r); i++ {
		assert.Contains(t, ArabicNumerals, string(r[i]))
	}
}

func TestRandString(t *testing.T) {
	r := RandString(10)
	for i := 0; i < len(r); i++ {
		assert.Contains(t, Alphabet, string(r[i]))
	}
}

func TestRandUpperString(t *testing.T) {
	r := RandUpperString(10)
	for i := 0; i < len(r); i++ {
		assert.Contains(t, UpperCharacter, string(r[i]))
	}
}

func TestRandLowerString(t *testing.T) {
	r := RandLowerString(10)
	for i := 0; i < len(r); i++ {
		assert.Contains(t, LowerCharacter, string(r[i]))
	}
}

func TestRandEnglishString(t *testing.T) {
	r := RandEnglishString(10)
	for i := 0; i < len(r); i++ {
		assert.Contains(t, EnglishString, string(r[i]))
	}
}

func TestRandBytes(t *testing.T) {
	randBytes := RandBytes(4)
	assert.Equal(t, 4, len(randBytes), "random bytes length must be equal to 4")
	assert.Equal(t, []byte{}, RandBytes(0),
		"null bytes slice must be equal to []byte{}",
	)
	assert.Equal(t, reflect.Slice, reflect.ValueOf(randBytes).Kind())

	// 正常情况测试用例
	// case 1
	b := RandBytes(3)
	if len(b) != 3 {
		t.Fatal("Data length is not three")
	}

	// case 2
	b = RandBytes(4)
	if len(b) != 4 {
		t.Fatal("Data length is not four")
	}

	// 异常测试用例
	// case 3
	b = RandBytes(-1)
	assert.Equal(t, b, nil)

}
