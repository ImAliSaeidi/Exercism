package romannumerals

import (
	"errors"
)

type RomanMap struct {
	Key   int
	Value string
}

var romanTable = []RomanMap{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("out of range")
	}

	result := ""
	for _, v := range romanTable {
		for input >= v.Key {
			result += v.Value
			input -= v.Key
		}
	}

	return result, nil
}
