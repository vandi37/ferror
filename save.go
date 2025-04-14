package ferror

import (
	"fmt"

	"github.com/vandi37/vanerrors"
)

type Save string

func (s Save) New(err error) error {
	return FError{
		Name: string(s),
		Err:  err,
	}
}

func (s Save) Errorf(format string, a ...any) error {
	return FError{
		Name: string(s),
		Err:  fmt.Errorf(format, a...),
	}
}

func (s Save) Wrap(message string, err error) error {
	return FError{
		Name: string(s),
		Err:  vanerrors.Wrap(message, err),
	}
}

func (s Save) Simple(message string) error {
	return FError{
		Name: string(s),
		Err:  vanerrors.Simple(message),
	}
}

func (s Save) Two(face, back error) error {
	return FError {
		Name: string(s),
		Err: TwoErrors(face, back),
	}
}

