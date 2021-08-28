package main

import "fmt"

func main() {
	var i int
	for i := 0; i < 10; i++ { // 서로 다른 i
		fmt.Println(i)
	}
	fmt.Println("최종 i의 값 : ", i)
}
