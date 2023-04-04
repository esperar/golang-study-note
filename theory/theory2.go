package theory

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i)
	}
	return total
}

func ForRange() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func IfElse() {
	fmt.Println(canIDrink(18))
	fmt.Println(canIDrink2(8))
}
