package util

import "golang.org/x/crypto/bcrypt"

func BcryptCompare(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func BcryptHash(password []byte, cost int) ([]byte, error) {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return nil, bcrypt.InvalidCostError(cost)
	}
	return bcrypt.GenerateFromPassword(password, cost)
}

func BcryptCost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}
