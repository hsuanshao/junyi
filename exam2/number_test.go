package exam2

import (
	"fmt"
	"testing"
)

func TestPrintNumber(t *testing.T) {
	testCase := []struct {
		Desc   string
		Input  int
		ExpRes string
		ExpErr error
	}{
		{
			Desc:   "normal case",
			Input:  15,
			ExpRes: "1, 2, 4, 7, 8, 11, 13, 14, 15",
			ExpErr: nil,
		},
	}

	for _, c := range testCase {
		res, err := printNumber(c.Input)

		if err != c.ExpErr {
			t.Error("expect error not equal", c.Desc)
		}

		if res != c.ExpRes {
			t.Error("expect result is not equal", c.Desc)
		}
		fmt.Println(res)
	}

}
