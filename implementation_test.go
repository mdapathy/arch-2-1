package lab1

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TestCase struct {
	name          string
	input         string
	expectedFloat float64
	expectedError error
}

var cases = []TestCase{
	// Error-free calculations
	{"Single operator", "5", 5, nil},
	{"Simple test 1", " - 3 8", -5, nil},
	{"Simple test 2", " + 3 / 3 8", 3.375, nil},
	{"Simple test 3", " + 3 / 3 8", 3.375, nil},
	{"Null division positive inf", "+ 4 / 5 0",math.Inf(1), nil },
/*	{"Complex test 1", " - 3 8", , nil},
	{"Complex test 2", " - 3 8", , nil},
	{"Complex test 3", " - 3 8", , nil},
	{"Complex test 4", " - 3 8", , nil},
	{"Complex test 5", " - 3 8", , nil},*/

	// Calculations returning error

	// regex-checked
	{"Empty string", "", 0, errors.New("the input cannot be treated as a Polish notation")},
	{"Delimiters string  ", "   ", 0,errors.New("the input cannot be treated as a Polish notation") },

	/*{"Improper formatting ", "5  *4 3", 0},
	{"Postfix notation", " 5 4 * 3 + ",0},
	{"Infix notation", " 6 ^ 4 + 3 *2", 0},
	{"Improper math operations", " 6 %5 - 4", 0},
	{"Not enough operations",  " 6 ^ 7 - 9 3.2", 0},
	{"Not enough operands",  " ^ - 5 + 3 4", 0},
	{"Improper float representation ", " + 5.4.5  * 4 3", 0},
	{"Single operand", " - ", 0, },*/

}

func TestPrefixEvaluation(t *testing.T) {

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			res, err := PrefixEvaluation(tcase.input)
			assert.Equal(t, tcase.expectedError, err)
			assert.Equal(t, tcase.expectedFloat, res)

		})
	}

}

func ExamplePrefixEvaluation() {
	//res, _ := PrefixEvaluation("/ + 2 3 2")
	//fmt.Println(res)

	// Output:
	// 2.5
}
