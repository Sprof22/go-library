package calculator

import (
	"fmt"
	"math"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) (int, error) {
	//handle case where y is zero
	if y == 0 {
		return 0, fmt.Errorf("division by zero yields infinity, which is an operation not allowed in simple arithmetic")
	}
	return x / y, nil
}

func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of a negative number is not allowed")
	}
	return math.Sqrt(a), nil
}

func PerimeterOfSquare(side float64) float64 {
	return 4 * side
}
