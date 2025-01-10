package main

import (
	"fmt"
	"go-library/array"
	"go-library/calculator"
)

func main() {
	//array operations
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)
	ascend, _ := array.SortAscending(arr)
	fmt.Println("Sorted Ascending:", ascend)
	descend, _ := array.SortDescending(arr)
	fmt.Println("Sorted Descending:", descend)
	min, _ := array.FindMin(arr)
	max, _ := array.FindMax(arr)
	fmt.Println("Minimum Value:", min)
	fmt.Println("Maximum Value:", max)

	//calculator operations
	fmt.Println("Addition:", calculator.Add(10, 20))
	fmt.Println("Subtraction:", calculator.Subtract(20, 10))
	fmt.Println("Multiplication:", calculator.Multiply(10, 20))
	division, _ := calculator.Divide(20, 10)
	fmt.Println("Division:", division)

	//area of a circle
	radius := 5.0
	fmt.Println("Area of Circle:", areaOfCircle(radius))
}
