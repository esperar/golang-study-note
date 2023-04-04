package main

import (
	"fmt"
	"github.com/esperer/learngo/something"
	"github.com/esperer/learngo/theory"
)

func main() {
	fmt.Println("Hello World!")

	something.SayHello()      // 패키지와 함수 호출 임포트
	theory.HelloName()        // 변수와 상수
	theory.FunctionsPartOne() // 함수 파트 1,2
	theory.ForRange()         // 반복문
}
