package calc

import "fmt"

func Add(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	fmt.Println("sum from package is", sum)
	return sum
}
