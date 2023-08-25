package roman_test

import (
	"go-chi-api/roman"
	"testing"
)

type TestCase struct {
	value    string
	expected int
	actual   int
}

func TestConvertRomanToInt(t *testing.T) {
	t.Run("should return 7 for VII", func(t *testing.T) {
		testCase := TestCase{
			value:    "VII",
			expected: 7,
		}

		testCase.actual = roman.ConvertRomanToInt(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
	t.Run("should return 9 for IX", func(t *testing.T) {
		testCase := TestCase{
			value:    "IX",
			expected: 9,
		}

		testCase.actual = roman.ConvertRomanToInt(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
