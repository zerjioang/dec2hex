package main

import (
	"dec2hex"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "hex-convert-5M-samples")
	total := 5000000
	var dst dec2hex.HexWrap // FFFFFFFFFFFFFFFF
	for i := 0; i < total; i++ {
		_ = dec2hex.FormatDst(&dst, uint64(i))
	}
}
