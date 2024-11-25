package util

import (
	"math/rand/v2"
	"testing"
)

func TestBcrypt(t *testing.T) {

	for i := 0; i < 100; i++ {
		randomPassword, _ := generateRandomBytes(rand.Int() % 72)
		randomCost := rand.Int()%10 + 4
		t.Logf("randomPassword: %s, randomCost: %d\n", randomPassword, randomCost)

		hashedPassword, err := BcryptHash(randomPassword, randomCost)
		t.Logf("hashedPassword: %s", hashedPassword)
		if err != nil {
			t.Errorf("bcrypt hash error: %v", err)
		}
		err = BcryptCompare(hashedPassword, randomPassword)
		if err != nil {
			t.Errorf("bcrypt compare error: %v", err)
		} else {
			t.Log("bcrypt compare success")
		}
		cost, err := BcryptCost(hashedPassword)
		if err != nil {
			t.Errorf("bcrypt cost error: %v", err)
		} else {
			t.Logf("bcrypt cost: %d", cost)
		}
	}
}
