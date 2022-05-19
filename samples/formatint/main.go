package main

import (
	"log"
	"strconv"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// WithFormatInt make uses of Go std FormatInt() to convert from
// decimal to hexadecimal
func WithFormatInt(v int64) string {
	return strconv.FormatInt(v, 16)
}

func main() {
	defer timeTrack(time.Now(), "hex-convert-5M-samples")
	total := 5000000
	for i := 0; i < total; i++ {
		_ = WithFormatInt(int64(i))
	}
}
