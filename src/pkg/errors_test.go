package pkg

import (
	"errors"
	"testing"
)

type DefinedError struct {
	message string
	code    int
}

func (de *DefinedError) Error() string {
	return de.message
}

func TestNewError(t *testing.T) {
	err := errors.New("test error")

	if err != nil {
		t.Log(err)
	}
}
