package helper

import (
	"log"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const RandStringLength = 10

var seedRando *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandStringBytes(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = LetterBytes[seedRando.Intn(len(LetterBytes))]
	}

	return string(b)
}

func WrappedError(errorMsg string, err error, logger *log.Logger) error {
	wrapperErr := errors.Wrap(err, errorMsg)
	logger.Println(wrapperErr.Error())

	return wrapperErr
}
