package util

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		v, _ := RandomInt(1, 100)
		t.Log(v)
	}
}

func TestRandomFloat(t *testing.T) {
	for i := 0; i < 100; i++ {
		v, _ := RandomFloat(1, 100)
		t.Log(v)
	}
}
