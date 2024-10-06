package errorx

import (
	"errors"
	"fmt"
)

func wrap(err error, s string, a ...any) error {
	return fmt.Errorf("%w: %s", err, fmt.Sprintf(s, a...))
}

func wrapFunc(err error) func(string, ...any) error {
	return func(s string, a ...any) error {
		return wrap(err, s, a...)
	}
}

var (
	ErrOK         = errors.New("OK")
	ErrConversion = errors.New("Conversion")
	ErrInvalid    = errors.New("Invalid")
	ErrNotFound   = errors.New("NotFound")
	ErrMarshal    = errors.New("Marshal")
	ErrUnmarshal  = errors.New("Unmarshal")
	ErrUnexpected = errors.New("Unexpected")
)

var (
	Conversion = wrapFunc(ErrConversion)
	Invalid    = wrapFunc(ErrInvalid)
	NotFound   = wrapFunc(ErrNotFound)
	Marshal    = wrapFunc(ErrMarshal)
	Unmarshal  = wrapFunc(ErrUnmarshal)
	Unexpected = wrapFunc(ErrUnexpected)
	OK         = wrapFunc(ErrOK)
)
