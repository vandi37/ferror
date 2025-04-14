package ferror

import (
	"errors"
	"fmt"
)

type two struct {
	face error
	back error
}

func (e two) Error() string {
	return fmt.Sprintf("%s: %s", e.face.Error(), e.back.Error())
}

func (e two) GetFace() error {
	return e.face
}

func (e two) GetBack() error {
	return e.back
}

func TwoErrors(face, back error) error {
	return two{face, back}
}

func GetFace(e error) error {
	var t two
	if !errors.As(e, &t) {
		return nil
	}
	return t.GetFace()
}

func GetBack(e error) error {
	var t two
	if !errors.As(e, &t) {
		return nil
	}
	return t.GetBack()
}

func (e two) Unwrap() []error {
	return []error{e.face, e.back}
}
