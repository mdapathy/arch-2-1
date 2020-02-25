package lab1

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO: document this function.
func PrefixEvaluation(input string) (float64, error) {

	// checks whether the line consists of valid math operations and integers/floats
	if validNotation, _ := regexp.Match(`^\s*((\+|-|\*|/|^)*\s+([0-9]+|[0-9]+\.*[0-9]+)\s*)+$`, []byte(input));
		!validNotation {
		return 0, errors.New("the input cannot be treated as a Polish notation")
	}

	var stack []float64

	partsOfString := strings.Fields(input)

	for i := len(partsOfString) - 1; i >= 0; i-- {

		if value, err := strconv.ParseFloat(partsOfString[i], 64); err == nil {
			stack = append(stack, value)
		} else if len(stack) >= 2 {
			value := evaluate(partsOfString[i], stack[len(stack)-1], stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, value)
		} else {
			return 0, errors.New("not enough operands in the input string")

		}
	}

	if len(stack) == 1 {
		return stack[0], nil
	} else {
		return 0, errors.New("not enough operations in the input string")

	}

}

func evaluate(operation string, first, second float64) float64 {
	switch operation {
	case "+":
		return first + second
	case "-":
		return first - second
	case "/":
		return first / second
	case "*":
		return first * second
	case "^":
		return math.Pow(first, second)
	default:
		return 0
	}
}
