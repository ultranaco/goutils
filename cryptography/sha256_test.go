package cryptography_test

import (
	"log"
	"testing"

	"github.com/ultranaco/goutils/cryptography"
)

func TestSha256String(t *testing.T) {
	hash := cryptography.Sha256String(PersonUltranaco{
		Name:     "Alejandro",
		Lastname: "Piña",
	})

	log.Println(hash)

	hash = cryptography.Sha256String(PersonUltranaco{
		Name:     "Alejandro",
		Lastname: "Piña Lindero",
	})

	log.Println(hash)

	hash = cryptography.Sha256String("")

	log.Println(hash)
}

type PersonUltranaco struct {
	Name     string
	Lastname string
}
