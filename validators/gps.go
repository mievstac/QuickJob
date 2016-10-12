package validators

import (
	"errors"

	multierror "github.com/hashicorp/go-multierror"
)

type GPS struct {
	X    float32
	Y    float32
	Z    float32
	A    float32
	Name string
}

func (g GPS) Validate() (bool, error) {
	var check bool = true
	var errorList error
	if g.X < 0 || g.Y < 0 {
		check = false
		errorList = multierror.Append(errorList, errors.New("GPS Check out of bounds"))
	}
	return check, errorList
}
