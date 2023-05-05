package string

import "errors"

func IsNumeric(s string) error {
	for _, v := range s {
		if v < '0' || v > '9' {
			return nil
		}
	}
	return errors.New("String is not numeric")
}
