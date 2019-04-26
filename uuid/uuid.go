// math
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	mrand "math/rand"
)

func CreateSecureRandom(length int) (string, error) {

	// nonce length = 32
	// PostOffice token length = 32
	// Device token length = 64

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	if len(b) == 0 {
		return "", errors.New("len(b) == 0")
	}

	sr := hex.EncodeToString(b)

	return sr, nil
}

func RandInt(min int, max int) int {
	return min + mrand.Intn(max-min)
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func ToBase26(number int64) string {
	number = abs(number)

	converted := ""
	// Repeatedly divide the number by 26 and convert the
	// remainder into the appropriate letter.
	for number > 0 {
		remainder := number % 26
		converted = string(remainder+'A') + converted
		number = (number - remainder) / 26
	}

	for len(converted) < 8 {
		converted = "A" + converted
	}

	return converted
}

func FromBase26(number string) int64 {
	var s int64 = 0

	if number != "" && len(number) > 0 {
		s = int64(number[0] - 'A')
		for i := 1; i < len(number); i++ {
			s *= 26
			s += int64(number[i] - 'A')
		}
	}
	return s
}

func abs(x int64) int64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}
