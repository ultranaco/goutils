package cryptography_test

import (
	"log"
	"testing"

	"github.com/ultranaco/goutils/cryptography"
)

func TestShortID(t *testing.T) {
	for index := 0; index < 500; index++ {
		shortID := cryptography.ShortID()
		log.Println(shortID)
	}
}
