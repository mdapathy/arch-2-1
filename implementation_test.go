package lab1

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

type TestCase struct {
	name          string
	input         string
	expectedFloat float64
	expectedError error
}

const equalityThreshold = 1e-9

var cases = []TestCase{
	// Error-free calculations
	{"Single operator", " 5 ", 5, nil},
	{"Simple test 1", " - 3 8", -5, nil},
	{"Simple test 2", " + 3 / 3 8", 3.375, nil},
	{"Simple test 3", " ^ 2 / 6 2", 8, nil},
	{"Null division positive inf", "+ 4 / 5 0", math.Inf(1), nil},
	{"Complex test 1", " - + * ^ 3 4  5 7 + * 6 7 8 ", 362, nil},                 // ( 3 ^ 4 ) * 5 + 7 - ( 6 * 7 + 8)
	{"Complex test 2", "- + 4.5 / * 6.7 8.9 6.7 +  6.6  * 4 3", -5.2, nil}, //  6.7 * 8.9 / 6.7 + 4.5  - ( 4 * 3 + 6.6)
	{"Complex test 3", " -  ^  * 3 4 6  / / * 4 9.3 3.7 0", math.Inf(-1), nil},   // ( 3 * 4 ) ^ 6 - 4 * 9.3 / 3.7 / 0
	{"Complex test 4", "- 67  + ^ / 3 4 4 / * 5 7 8 ", 62.30859375, nil}, //  67 - ((3 / 4) ^ 4 + (5 * 7 /8) )

	// Calculations returning error
	// regex-checked
	{"Empty string", "", 0, errors.New("input string is not in Polish Notation")},
	{"Whitespace string  ", "	", 0, errors.New("input string is not in Polish Notation")},
	{"Improper formatting ", "5  *4 3", 0, errors.New("input string is not in Polish Notation")},
	{"Postfix notation ", "3  4 + 7 *", 0, errors.New("input string is not in Polish Notation")},
	{"Improper math operations", " 6 % 5 - 4", 0, errors.New("input string is not in Polish Notation")},
	{"Improper float representation ", " + 5.4.5  * 4 3", 0, errors.New("input string is not in Polish Notation")},
	{"Single operand", " - ", 0, errors.New("input string is not in Polish Notation")},

	//stack evaluation based
	{"Not enough operations", " 6 ^ 7 - 9 3.2", 0, errors.New("not enough operations in the input string")},
	{"Not enough operands", " ^ - 5 + 3 4", 0, errors.New("not enough operands in the input string")},
	{"Infix notation", " 6 ^ 4 + 3 * 2", 0, errors.New("not enough operands in the input string")},
}

func TestPrefixEvaluation(t *testing.T) {

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			res, err := PrefixEvaluation(tcase.input)
			require.Equal(t, tcase.expectedError, err)
			assert.True(t, tcase.expectedFloat == res || math.Abs(tcase.expectedFloat - res) < equalityThreshold)

		})
	}

}

func ExamplePrefixEvaluation() {
	res, err := PrefixEvaluation(" * 4 / + 2 3 2")
	if err == nil {
		fmt.Println(res)
	}

	// Output:
	// 10
}
