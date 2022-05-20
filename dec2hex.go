package dec2hex

type HexWrap [16]uint8

type Hex []byte

// String converts Hex ([]byte) representation to string
func (h Hex) String() string {
	return string(h)
}

// Format converts uint64 to Hex value
// using an internal byte destination slice
func Format(n uint64) Hex {
	var dst HexWrap // FFFFFFFFFFFFFFFF
	return FormatDst(&dst, n)
}

// FormatDst converts uint64 to Hex value
// using provided byte destination slice
func FormatDst(dst *HexWrap, n uint64) Hex {
	if n == 0 {
		return Hex("0")
	}
	var idx uint8 = 16
	_ = dst[15]
	for q := n; q > 0; q >>= 4 {
		m := q % 16
		dst[idx-1] = uint8(48 + m)
		if m > 9 {
			dst[idx-1] += 7
		}
		idx--
	}
	return dst[idx:]
}
