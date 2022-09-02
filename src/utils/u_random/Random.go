package u_random

import "math/rand"

const (
	SEED_NUMBER    string = "1234567890"
	SEED_LOWER     string = "qwertyuiopasdfghjklzxcvbnm"
	SEED_UPPER     string = "QWERTYUIOPASDFGHJKLZXCVBNM"
	SEED_CHARACTER string = "~!@#$%^&*()_+`-=[]{}|;:,./<>?"
)

func GenString(seeds string, length int) string {
	if len(seeds) == 0 && length <= 0 {
		return ``
	}
	seedsBytes := []byte(seeds)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, seedsBytes[rand.Intn(len(seedsBytes))])
	}
	return string(result)
}
