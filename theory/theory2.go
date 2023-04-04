package theory

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func ForRange() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
