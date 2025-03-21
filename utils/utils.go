package utils

import (
	"errors"
	"strconv"
)

func ToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("não é permitido caractere diferente de número. por favor digite um número")
	}

	return i, nil
}
