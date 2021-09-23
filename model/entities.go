package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type  string
	Size  int
	Value string
}

func NewResult(input string) (Result, error) {
	return createResult(input)
}

func createResult(input string) (Result, error) {
	size, err := getSize(input)
	if err == nil {
		value, e := getValue(input, size)
		if e == nil {
			return Result{getType(input), size, value}, nil
		}
		return Result{}, e
	}
	return Result{}, err
}

func getType(input string) string {
	return input[0:2]
}

func getSize(input string) (int, error) {
	size, err := strconv.Atoi(input[2:4])
	if err == nil {
		return size, nil
	}
	return 0, errors.New("value must be numeric")
}

func getValue(input string, lenght int) (string, error) {
	realSize := len(input[0:])
	endOfInput := lenght + 4
	if endOfInput <= realSize {
		return input[4:endOfInput], nil
	}
	return "", errors.New("incorrect lenght value: index out of range")
}
