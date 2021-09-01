package main

import (
	"fmt"
)

func main() {
	// 동적 배열 make(type, 초기값, capacity)
	var a []int
	b := []int{1, 2, 3}
	c := make([]int, 3)
	d := make([]int, 0, 8)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))
	fmt.Printf("len(c) = %d\n", len(c))
	fmt.Printf("cap(c) = %d\n", cap(c))
	fmt.Printf("len(d) = %d\n", len(d))
	fmt.Printf("cap(d) = %d\n", cap(d))

	a = append(a, 1)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	x1 := []int{1, 2}
	y1 := append(x1, 1)

	fmt.Printf("%p %p\n", x1, y1) // %p :주소

	x2 := make([]int, 2, 4)
	y2 := append(x2, 1)

	fmt.Printf("%p %p\n", x2, y2)

	x3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y3 := x3[4:8] // 5, 6, 7, 8
	z3 := x3[4:]
	// y3[0] = 1
	// y3[1] = 2 // slice는 x3에서 가리키는 곳이 달라지는것!
	// slice하여 다른 변수로 지정하여 값을 변경하면 본래의 배열의 값도 변함! 따라서 복사해서 사용해야 문제 해결

	fmt.Println(y3)
	fmt.Println(z3)

	for i := 0; i < 5; i++ {
		// var back int
		var front int
		// x3, back = RemoveBack(x3)
		x3, front = RemoveFront(x3)
		// fmt.Printf("%d, ", back)
		fmt.Printf("%d, ", front)
	}
	fmt.Println(x3)
}

// slice라는건 나머지 부분을 없애는 것이 아니고, 가리키는 곳을 변경하는것! len, cap

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
