package log

import "crypto/rand"

const (
	upperCaseAlphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// getEpochRandomString generates a random string with the provided length using the given alphabet
func getEpochRandomString() (string, error) {
	randomBytes := make([]byte, 5)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	for index, randomByte := range randomBytes {
		foldedOffset := randomByte % byte(len(upperCaseAlphanumeric))
		randomBytes[index] = upperCaseAlphanumeric[foldedOffset]
	}
	return string(randomBytes), nil
}
