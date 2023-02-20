package log

import (
	"log"

	"github.com/pkg/errors"
)

func Error(err error, message string) {
	newErr := errors.Wrap(err, message)
	log.Println(newErr)
}

func Info(message string) {
	log.Println(message)
}
