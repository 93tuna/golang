package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func fun1(x, y int) (int, int) {
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println(x, y)
}

func fun3(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("fun3(%d) before call fun3(%d)\n", x, x-1)
	fun3(x - 1)
	fmt.Printf("fun3(%d) after call fun3(%d)\n", x, x-1)
}

func sum(x, y int) int {
	if x == 0 {
		return y
	}
	y += x
	return sum(x-1, y)
}

func main() {
	/* for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a, b := fun1(2, 3)

	fmt.Println(a, b) */

	fun3(10)

	fmt.Println(sum(25, 0))

}
