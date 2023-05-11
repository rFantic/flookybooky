package test

import (
	"flookybooky/internal/validate"
	"fmt"
)

func Example_numeric() {
	res := validate.IsNumeric("0123456789")
	fmt.Println(res)
	// Output:
	// <nil>
}
