package calc

func Add(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}
