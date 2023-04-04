package theory

import "fmt"

func HelloName() {
	const name = "hope"
	// name = "kim" Error
	fmt.Println(name)

	var name1 string = "kim"
	name1 = "hope2"
	fmt.Println(name1)
}
