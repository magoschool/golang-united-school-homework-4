package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

const errorMessage = "StringSum failed: %w"

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func skipSpace(input string, index *int) {
	for i := *index; i < len(input); i++ {
		if input[i] > ' ' {
			*index = i
			return
		}
	}
	*index = len(input)
}

func tryGetSign(aInput string, index *int, aSignValue *byte) bool {
	skipSpace(aInput, index)
	i := *index
	if i < len(aInput) && (aInput[i] == '-' || aInput[i] == '+') {
		*aSignValue = aInput[i]
		*index = i + 1
		skipSpace(aInput, index)
		return true
	}
	return false
}

func isDigit(aChar byte) bool {
	return aChar >= '0' && aChar <= '9'
}

func tryGetValue(aInput string, aIndex *int, aValue *int) bool {
	skipSpace(aInput, aIndex)
	i := *aIndex
	for i < len(aInput) && isDigit(aInput[i]) {
		i++
	}

	if i > *aIndex {
		var lNumber = aInput[*aIndex:i]
		lValue, lErr := strconv.Atoi(lNumber)
		if lErr != nil {
			return false
		}

		*aValue = lValue
		*aIndex = i
		skipSpace(aInput, aIndex)
		return true
	}

	return false
}

func StringSum(input string) (output string, err error) {
	var lIndex int = 0
	var lSign byte = 0
	var lValue1 int = 0
	var lValue2 int = 0

	skipSpace(input, &lIndex)
	if lIndex >= len(input) {
		return "", fmt.Errorf(errorMessage, errorEmptyInput)
	}

	if !tryGetSign(input, &lIndex, &lSign) {
		lSign = '+'
	}
	if !tryGetValue(input, &lIndex, &lValue1) {
		return "", fmt.Errorf(errorMessage, errorNotTwoOperands)
	}
	if lSign == '-' {
		lValue1 = -lValue1
	}

	if !tryGetSign(input, &lIndex, &lSign) {
		return "", fmt.Errorf(errorMessage, errorNotTwoOperands)
	}
	var lSign2 byte = 0
	if tryGetSign(input, &lIndex, &lSign2) {
		if lSign2 == '-' {
			if lSign == '+' {
				lSign = '-'
			} else {
				lSign = '+'
			}
		}
	}

	if !tryGetValue(input, &lIndex, &lValue2) {
		return "", fmt.Errorf(errorMessage, errorNotTwoOperands)
	}
	if lSign == '-' {
		lValue2 = -lValue2
	}

	skipSpace(input, &lIndex)
	if lIndex < len(input) {
		return "", fmt.Errorf(errorMessage, errorNotTwoOperands)
	}

	return strconv.Itoa(lValue1 + lValue2), nil
}
