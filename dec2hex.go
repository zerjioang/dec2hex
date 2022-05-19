package dec2hex

type HexWrap [16]byte

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
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 {
			// write hexadecimal alpha byte
			dst[idx-1] = byte('7' + m)
		} else {
			// write hexadecimal numeric byte
			dst[idx-1] = byte('0' + m)
		}
		idx--
	}
	return dst[idx:]
}
