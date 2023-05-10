package test

import (
	"flookybooky/internal/string"
	"fmt"
)

func Example_numeric() {
	res := string.IsNumeric("0123456789")
	fmt.Println(res)
}
