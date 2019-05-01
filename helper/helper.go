package helper

import (
	"github.com/pkg/errors"
	"log"
	"math/rand"
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const RandStringLength = 10

func RandStringBytes(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}
	return string(b)
}

func WrappedError(errorMsg string, err error, logger *log.Logger)  error{
	wrapperErr := errors.Wrap(err, errorMsg)
	logger.Println(wrapperErr.Error())
	return wrapperErr
}
