package main

func main() {
	a := 2
	b := &a
	a = 5
	println(*b)
}
