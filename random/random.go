// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/3/9 21:42
// @Author      : Createitv
// @FileName    : random.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

// Package random implements some function to generate random Letters or ASCIICharacters
package random

import (
	craned "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	Digits                string = "0123456789"                                                         // Digits: [0-9]
	ASCIILettersLowercase string = "abcdefghijklmnopqrstuvwxyz"                                         // Ascii Lowerrcase Letters: [a-z]
	ASCIILettersUppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"                                         // Ascii Uppercase Letters: [A-Z]
	Letters               string = ASCIILettersLowercase + ASCIILettersUppercase                        // Ascii Letters: [a-zA-Z]
	ASCIICharacters       string = Letters + Digits                                                     // Ascii Charaters: [a-zA-Z0-9]
	Hexdigits             string = "0123456789abcdefABCDEF"                                             // Hex Digits: [0-9a-fA-F]
	Octdigits             string = "01234567"                                                           // Octal Digits: [0-7]
	Punctuation           string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                                 // Punctuation and special characters
	Printable             string = Digits + ASCIILettersLowercase + ASCIILettersUppercase + Punctuation // Printable

)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandBetween 生成min到max之间的随机数，包括min不包括max
func RandBetween(min, max int) int {
	if max == min {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

// RandBytes 生成 length 长度的字节切片
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)
	_, _ = io.ReadFull(craned.Reader, b)

	return b
}

// RandFromString 从一段给定字符串中生成一定长度的字符串
func RandFromString(s string, length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}

	return string(b)
}

// RandNumber 生成 length 长度的只有数字的字符串
func RandNumber(length int) string {
	return RandFromString(Digits, length)
}

// RandString 生成 length 长度的只有52个大小写英文字母的字符串
func RandString(length int) string {
	return RandFromString(Letters, length)
}

// RandUpperString 生成 length 长度的只有26个大写英文字母的字符串
func RandUpperString(length int) string {
	return RandFromString(ASCIILettersUppercase, length)
}

// RandLowerString 生成 length 长度的只有26个小写英文字母的字符串
func RandLowerString(length int) string {
	return RandFromString(ASCIILettersLowercase, length)
}

// RandEnglishString 生成 length 长度的只有52个大小写英文字母 + 数字的字符串
func RandEnglishString(length int) string {
	return RandFromString(ASCIICharacters, length)
}

// UUIdV4 生成依据RFC 4122 的UUID
func UUIdV4() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(craned.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10],
		uuid[10:],
	), nil
}

// Choice makes a random choice from a slice of string.
func Choice(j []string) string {
	choose := RandBetween(0, len(j))

	return j[choose]
}

// RandHex 生成 length 长度的只有16进制数字的字符串
func RandHex(length int) string {
	return RandFromString(Hexdigits, length)
}

// RandOct 生成 length 长度的只有8进制数字的字符串
func RandOct(length int) string {
	return RandFromString(Octdigits, length)
}
