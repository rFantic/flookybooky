package test

import (
	"flookybooky/internal/string"
	"fmt"
)

func ExampleNumeric() {
	res := string.IsNumeric("0123456789")
	fmt.Println(res)
}
