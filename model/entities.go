package model

import (
	"errors"
	"regexp"
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

/*
	I didn't find other way to show the error, that's why I use many if - else
*/
func createResult(input string) (Result, error) {
	size, err := getSize(input)
	if err == nil {
		value, er := getValue(input, size)
		if er == nil {
			tp := getType(input)
			isCorrect, e := validateTypeValue(tp, value)
			if isCorrect {
				return Result{tp, size, value}, nil
			} else {
				return Result{}, e
			}
		} else {
			return Result{}, er
		}
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

func validateTypeValue(tp, value string) (bool, error) {
	switch tp {
	case "TX":
		return validateTypeTX(value)
	case "NN":
		return isNumeric(value)
	default:
		return false, errors.New("something went wrong")
	}
}

func validateTypeTX(value string) (bool, error) {
	re := regexp.MustCompile("[0-9]+")
	p := re.FindAllString(value, -1)
	if len(p) > 0 {
		return false, errors.New("type and value doesn't match")
	}
	return true, nil
}

func isNumeric(value string) (bool, error) {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false, errors.New("type and value doesn't match")
	}
	return true, nil
}
