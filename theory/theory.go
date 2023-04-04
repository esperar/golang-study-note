package theory

import (
	"fmt"
	"strings"
)

func HelloName() {
	const name = "hope"
	// name = "kim" Error
	fmt.Println(name)

	var name1 string = "kim"
	name1 = "hope2"
	fmt.Println(name1)
}

func FunctionsPartOne() {
	fmt.Println(multiply(2, 2))
	fmt.Println(lenAndUpper("kim hope"))
	totalLength, upperName := lenAndUpper("kim hope")
	// totalLength, _ 이런식으로 작성하면 값 하나를 무시하고 앞의 값만 받을 수 있다.
	fmt.Println(totalLength, upperName)
	repeatMe("eee", "aaa", "bbb", "ccc")
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
