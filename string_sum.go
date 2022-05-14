package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	var varOperand = 0

	strings.ReplaceAll(input, " ", "")
	if input == "" {
		err = fmt.Errorf("something went wrong: %w", errorEmptyInput)
		output = "0"
	} else {
		var resultExpression = 0
		result := strings.Split(input, "")
		var isMinus bool
		for _, element := range result {
			if element == "+" {
				isMinus = false
			} else if element == "-" {
				isMinus = true
			} else {
				varOperand++
				value, _ := strconv.Atoi(element)
				if isMinus {
					resultExpression -= value
				} else {
					resultExpression += value
				}
			}
		}
		if varOperand != 2 {
			err = fmt.Errorf("something went wrong: %w", errorNotTwoOperands)
		} else {
			output = fmt.Sprint(resultExpression)
		}
	}
	return output, err
}
