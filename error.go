package ferror

import (
	"errors"
	"fmt"

	"github.com/vandi37/vanerrors"
)

type FError struct {
	Name string
	Err  error
}

func (f FError) Error() string {
	return fmt.Sprintf("%s(%s)", f.Name, f.Err.Error())
}

func (f FError) Unwrap() error {
	return f.Err
}

func New(name string, err error) error {
	return FError{
		Name: name,
		Err:  err,
	}
}

func Errorf(name string, format string, a ...any) error {
	return FError{
		Name: name,
		Err:  fmt.Errorf(format, a...),
	}
}

func Wrap(name, message string, err error) error {
	return FError{
		Name: name,
		Err:  vanerrors.Wrap(message, err),
	}
}

func Simple(name, message string) error {
	return FError{
		Name: name,
		Err:  vanerrors.Simple(message),
	}
}

func Two(name string, face, back error) error {
	return FError{
		Name: name,
		Err:  TwoErrors(face, back),
	}
}

func (f FError) GetName() string {
	return f.Name
}

func GetName(err error) (string, bool) {
	var f FError
	ok := errors.As(err, &f)
	if ok {
		return f.Name, true
	} else {
		return "", false
	}
}
