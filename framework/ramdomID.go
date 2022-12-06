package framework

import (
	"crypto/rand"

	"github.com/google/uuid"
)

type RandomID interface {
	Generate(int) string
	GenerateUUID4() (string, error)
}

type ramdomIDGeneratorImpl struct {
}

func NewRamdomIDGenerator() RandomID {
	return &ramdomIDGeneratorImpl{}
}

func (r *ramdomIDGeneratorImpl) Generate(length int) string {
	const chars = "ABCDEFGHIJKLMNPQRSTWXYZ123456789"
	b := make([]byte, length)

	rand.Read(b)

	var result string
	for _, v := range b {
		result += string(chars[int(v)%len(chars)])
	}

	return result
}

func (r *ramdomIDGeneratorImpl) GenerateUUID4() (string, error) {
	uuidObj, err := uuid.NewRandom()
	return uuidObj.String(), err
}
