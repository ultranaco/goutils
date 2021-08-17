package cryptography

import (
	"crypto/sha256"
	"fmt"
)

func Sha256String(object interface{}) string {

	hex := sha256.New()
	rawObject := fmt.Sprintf("%+v", object)
	hex.Write([]byte(rawObject))

	return fmt.Sprintf("%x", hex.Sum(nil))
}
