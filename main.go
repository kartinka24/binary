package main

import (
	"fmt"

	op "github.com/binary/arithmetic"
)

func main() {
	three1 := []int{2, 1, 2}
	three2 := []int{1, 0, 1, 1}
	three3 := []int{2, 2}

	fmt.Printf("%v + %v = %v\n", three1, three2, op.Sum(three1, three2, 3))
	fmt.Printf("%v - %v = %v\n", three1, three2, op.Difference(three1, three2, 3))
	fmt.Printf("%v * %v = %v\n", three1, three2, op.Multiply(three1, three2, 3))
	fmt.Printf("%v + %v = %v\n", three1, three3, op.Sum(three1, three3, 3))
}
