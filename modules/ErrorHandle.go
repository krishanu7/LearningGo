package modules

import "errors"

func Division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return x / y, nil
	}
}