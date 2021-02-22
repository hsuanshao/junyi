package exam1

import "testing"

func TestRevertString(t *testing.T) {
	testCase := []struct {
		Desc   string
		Input  string
		ExpRes string
	}{
		{
			Desc:   "normal case",
			Input:  "junyiacademy",
			ExpRes: "ymedacaiynuj",
		},
	}

	for _, c := range testCase {
		result := revertString(c.Input)

		if result != c.ExpRes {
			t.Error("revert failure", c.Desc)
			continue
		}
	}
}

func TestRevertSentence(t *testing.T) {
	testCase := []struct {
		Desc   string
		Input  string
		ExpRes string
	}{
		{
			Desc:   "normal sentence case",
			Input:  "flipped class room is important",
			ExpRes: "deppilf ssalc moor si tnatropmi",
		},
	}

	for _, c := range testCase {
		result := revertSentence(c.Input)

		if result != c.ExpRes {
			t.Error("revert sentence failure", c.Desc)
			continue
		}
	}
}
