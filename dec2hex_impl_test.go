// implementation of test methods for performance comparison

package dec2hex

import (
	"fmt"
	"log"
	"strconv"
)

// WithFormatInt make uses of Go std FormatInt() to convert from
// decimal to hexadecimal
func WithFormatInt(v int64) string {
	return strconv.FormatInt(v, 16)
}

// WithFmt make uses of Go std fmt.Sprintf() to convert from
// decimal to hexadecimal
func WithFmt(v int64) string {
	return fmt.Sprintf("%x", v)
}

// WithGithubCode1 makes uses of github algorithm at
// https://github.com/TDCQZD/GoDev/blob/982ca325329baab67489d7f3d4f43dcdc8bdf561/utils/convert.go#L53
func WithGithubCode1(v int64) string {
	if v < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if v == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := v; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			// ./dec2hex_impl_test.go:39:28: conversion from int64 to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
			m2 := string(rune(m))
			s = fmt.Sprintf("%v%v", m2, s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}
