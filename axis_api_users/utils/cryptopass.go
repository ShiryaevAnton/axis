package utils

import (
	"crypto/sha256"
	"fmt"
)

//GetSHA256 ...
func GetSHA256(input string) string {
	sum := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", sum)
}
