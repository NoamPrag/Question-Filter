package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(_ *testing.T) {
	x, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	} else {
		fmt.Println("Number:")
		fmt.Println(x)
	}
}
