package cryptography

import (
	"fmt"
	"time"
)

var alphabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func ShortID() string {
	var base62Hash string

	length := int64(len(alphabet))
	now := time.Now()
	nanos := now.UnixNano() / 100
	offsetIndex := nanos % 10

	base62Hash = alphabet[nanos%length]
	nanos = nanos / length

	for nanos > 0 {
		indexAlphabet := nanos % length
		base62Hash = alphabet[indexAlphabet] + base62Hash
		nanos = nanos / length
	}

	nose := base62Hash[offsetIndex:]
	tail := base62Hash[:offsetIndex]

	return fmt.Sprintf("%s%s%v", nose, tail, alphabet[offsetIndex])
}
