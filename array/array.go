//⁠Sort numbers of a given array (both ascending and descending order)

package array

import (
	"fmt"
)

func SortAscending(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("array is empty")
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr, nil
}

func SortDescending(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("array is empty")
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr, nil
}

func FindMin(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min, nil
}

func FindMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max, nil
}
