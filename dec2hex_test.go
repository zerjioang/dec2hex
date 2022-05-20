package dec2hex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/easypprof"
	"math"
	"testing"
)

const (
	maxInt64 int64 = math.MaxInt64
)

func TestFormat(t *testing.T) {
	t.Run("format-int", func(t *testing.T) {
		// max
		assert.Equal(t, WithFormatInt(maxInt64), "7fffffffffffffff")
		// min
		assert.Equal(t, WithFormatInt(0), "0")
	})
	t.Run("fmt", func(t *testing.T) {
		// max
		assert.Equal(t, WithFmt(maxInt64), "7fffffffffffffff")
		// min
		assert.Equal(t, WithFmt(0), "0")
	})
	t.Run("github-1", func(t *testing.T) {
		// max
		assert.Equal(t, WithGithubCode1(maxInt64), "7FFFFFFFFFFFFFFF")
		// min
		assert.Equal(t, WithGithubCode1(0), "0")
	})
	t.Run("ours-format", func(t *testing.T) {
		// max
		assert.Equal(t, Format(uint64(maxInt64)).String(), "7FFFFFFFFFFFFFFF")
		// min
		assert.Equal(t, Format(0).String(), "0")
	})
	t.Run("ours-format-dst", func(t *testing.T) {
		// max
		var dst HexWrap
		assert.Equal(t, FormatDst(&dst, uint64(maxInt64)).String(), "7FFFFFFFFFFFFFFF")
		// min
		assert.Equal(t, FormatDst(&dst, 0).String(), "0")
	})
	t.Run("profile", func(t *testing.T) {
		if false {
			var dst HexWrap
			easypprof.Profile(t, 500000000, func() {
				FormatDst(&dst, uint64(maxInt64))
			})
		}
	})
	t.Run("shift", func(t *testing.T) {
		var n = 512
		fmt.Println(n)
		fmt.Println(512 >> 4)
		assert.Equal(t, 512>>4, 32)
	})
}

func BenchmarkFormat(b *testing.B) {
	b.Run("format-int", func(b *testing.B) {
		var v string
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = WithFormatInt(maxInt64)
		}
		assert.Equal(b, v, "7fffffffffffffff")
	})
	b.Run("fmt", func(b *testing.B) {
		var v string
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = WithFmt(maxInt64)
		}
		assert.Equal(b, v, "7fffffffffffffff")
	})
	b.Run("github-1", func(b *testing.B) {
		var v string
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = WithGithubCode1(maxInt64)
		}
		assert.Equal(b, v, "7FFFFFFFFFFFFFFF")
	})
	b.Run("ours-format", func(b *testing.B) {
		var v Hex
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = Format(uint64(maxInt64))
		}
		assert.Equal(b, v.String(), "7FFFFFFFFFFFFFFF")
	})
	b.Run("ours-format-dst", func(b *testing.B) {
		var v Hex
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		var dst HexWrap
		for i := 0; i < b.N; i++ {
			v = FormatDst(&dst, uint64(maxInt64))
		}
		assert.Equal(b, v.String(), "7FFFFFFFFFFFFFFF")
	})
}
