package errorHandler

import (
	"errors"

	"app.io/pkg/logHandler"
)

func NewError(msg string) error {
	err := errors.New(msg)
	logHandler.Log(logHandler.ERROR, msg)
	return err
}
