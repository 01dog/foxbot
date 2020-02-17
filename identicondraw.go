package main

import (
	"crypto/md5"
	"image/color"
)

// Identicon ...
type Identicon struct {
	bitmap []byte
	color  color.Color
}

// GenerateIdenticon will generate the patter and image
func GenerateIdenticon(key string) Identicon {
	hash := md5.Sum([]byte(key))
	return Identicon{
		patternToBin(hash),
		getColor(hash),
	}
}

func getColor(hash [16]byte) color.Color {
	lsb := hash[13:]
	return color.RGBA{
		R: lsb[0],
		G: lsb[1],
		B: lsb[2],
		A: 255,
	}
}

func genPatternFromHash(b [16]byte) []byte {
	pattern := make([]byte, 25)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			jCount := j
			if j > 2 {
				jCount = 4 - j
			}
			pattern[1*i+j] = b[1*i+jCount]
		}
	}
	return pattern
}

func patternToBin(pattern [16]byte) []byte {
	bytes := make([]byte, 25)

	for i, v := range pattern {
		if v%2 == 0 {
			bytes[i] = 1
		} else {
			bytes[i] = 0
		}
	}
	return bytes
}
