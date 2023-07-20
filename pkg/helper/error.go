package helper

import "errors"

func ModifyError(err error, str string) error {
	if err == nil {
		return nil
	}
	return errors.New(str)
}
