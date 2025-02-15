package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	var varOperand = 0
	if input == "" {
		return "", GetErrorEmptyInput()
	} else if len(input) < 3 {
		return "", GetErrorNotTwoOperands()
	} else {
		var resultExpression = 0
		var isMinus bool
		var indexCurrent int

		inputValue := []rune(input)

		for index, element := range inputValue {
			if element == '+' {
				isMinus = false
			} else if element == '-' {
				isMinus = true
			} else if indexCurrent >= index && indexCurrent != 0 {
			} else {
				var digit []rune
				digit = append(digit, element)
				for index < len(inputValue)-1 && (unicode.IsDigit(inputValue[index+1]) ||
					unicode.IsLetter(inputValue[index+1])) {
					digit = append(digit, inputValue[index+1])
					index++
					indexCurrent = index
				}
				value, er := strconv.Atoi(string(digit))
				if er != nil {
					return "", fmt.Errorf("%w", er)
				}
				varOperand++

				if isMinus {
					resultExpression -= value
				} else {
					resultExpression += value
				}
			}
		}
		if varOperand != 2 {
			return "", GetErrorNotTwoOperands()
		} else {
			output = fmt.Sprint(resultExpression)
		}
	}

	return output, err
}

func GetErrorEmptyInput() error {
	return fmt.Errorf("%w", errorEmptyInput)
}

func GetErrorNotTwoOperands() error {
	return fmt.Errorf("%w", errorNotTwoOperands)
}
