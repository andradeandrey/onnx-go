package testbackend

import (
	"fmt"
	"testing"
)

var testCases = []TestCase{
	TestCase{
		Title: "case1",
	},
	TestCase{
		Title: "case2",
	},
	TestCase{
		Title: "case3",
	},
	TestCase{
		Title: "case4",
	},
}

func TestOperators(t *testing.T) {

	var i int
	var tc TestCase
	for i, tc = range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("%s", tc.Title), func(t *testing.T) {
			t.Parallel()
			t.Fail()
		})
	}
	t.Log(i)
}
