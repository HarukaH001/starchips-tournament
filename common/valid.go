package common

import "github.com/go-playground/validator/v10"

func Valid(s any) error {
	if err := validator.New().Struct(s); err != nil {
		return err
	}
	return nil
}

func MustValid(s any) {
	if err := Valid(s); err != nil {
		panic(err)
	}
}
