package random_creation_service

import (
	"math/rand"
	"time"
)

type RandomCreationService struct{}

func NewRandomCreationService() *RandomCreationService {
	return &RandomCreationService{}
}

func (randomService *RandomCreationService) GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())

	digits := "0123456789"
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		otp[i] = digits[rand.Intn(len(digits))]
	}

	return string(otp)
}
