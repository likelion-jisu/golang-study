package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// 중복된 이름의 두 패키지를 가져올 경우
// 패키지 이름은 섀도잉될 수 있다.

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic(err)
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
