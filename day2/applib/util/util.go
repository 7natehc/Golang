package util

import (
	"errors"
)

func square(num int) int {
	return num * num
}

func divide(num1, num2 float32) (float32, error) {
	//check if num2==0
	if num2 == 0 {
		return 0, errors.New("type")
	}
	return (num1 / num2), nil
}
