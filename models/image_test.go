package models

import "testing"

type NewImageArgs struct {
	X, Y, Zoom string
}

type TestCase struct {
	Args     NewImageArgs
	Expected Image
}

var testCases = []TestCase{
	{
		Args:     NewImageArgs{"0", "0", "1"},
		Expected: Image{X: 0, Y: 0, Zoom: 1},
	},
}

func TestNewImage(t *testing.T) {
	for _, testCase := range testCases {
		parsedValue, err := New("test.png", testCase.Args.X, testCase.Args.Y, testCase.Args.Zoom)

		if err != nil {
			t.Error(err)
		}

		if parsedValue.X != testCase.Expected.X || parsedValue.Y != testCase.Expected.Y || parsedValue.Zoom != testCase.Expected.Zoom {
			t.Error("for args: ",
				testCase.Args,
				"expected: ",
				testCase.Expected,
				" ,but got: ",
				parsedValue,
			)
		}
	}
}
