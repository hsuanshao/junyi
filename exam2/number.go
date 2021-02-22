package exam2

import (
	"errors"
	"strconv"
)

func printNumber(input int) (output string, err error) {
	if input < 1 {
		return "", errors.New("number must large than 1")
	}

	for i := 1; i <= input; i++ {
		threePle := false
		fivePle := false

		if i%3 == 0 {
			threePle = true
		}

		if i%5 == 0 {
			fivePle = true
		}

		if threePle == false && fivePle == true {
			continue
		}

		if threePle == true && fivePle == false {
			continue
		}

		output += strconv.FormatInt(int64(i), 10)
		if i < input {
			output += ", "
		}
	}

	return output, nil
}
