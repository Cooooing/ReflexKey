package util

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
)

func RandomInt(min int64, max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

func RandomFloat(min float64, max float64) (float64, error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 0, err
	}
	u := binary.LittleEndian.Uint64(b[:])
	f := math.Nextafter(float64(u)/math.MaxUint64, 1)
	result := min + f*(max-min)
	if result >= max {
		result = max - math.SmallestNonzeroFloat64
	}
	return result, nil
}
