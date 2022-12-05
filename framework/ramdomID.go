package framework

import (
	"crypto/rand"
	"github.com/google/uuid"
)

type RamdomIDGenerator interface {
	Generate(int) string
	GenerateUUID4() (string, error)
}

type RamdomIDGeneratorImpl struct {
}

func NewRamdomIDGenerator() RamdomIDGenerator {
	return &RamdomIDGeneratorImpl{}
}

func (r *RamdomIDGeneratorImpl) Generate(length int) string {
	const chars = "ABCDEFGHIJKLMNPQRSTWXYZ123456789"
	b := make([]byte, length)

	rand.Read(b)

	var result string
	for _, v := range b {
		result += string(chars[int(v)%len(chars)])
	}

	return result
}

func (r *RamdomIDGeneratorImpl) GenerateUUID4() (string, error) {
	uuidObj, err := uuid.NewRandom()
	return uuidObj.String(), err
}
