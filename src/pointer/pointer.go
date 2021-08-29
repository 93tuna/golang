package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int

	grade   string
	subject string
}

func (s Student) ViewSungJuk() { // 포인터 형태로 인자를 받으면 메모리 주소만 복사, 값으로 받으면 전체 속성 복사
	fmt.Println(s.subject, s.grade)
}

func (s *Student) InputSungJuck(subject string, grade string) {
	s.subject = subject
	s.grade = grade
}

func main() {
	a := 3
	b := 3
	var p *int

	p = &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)

	x := 1

	Increase(&x)

	fmt.Println(x)

	// var s Student
	s := Student{name: "dong", age: 29, subject: "수학", grade: "A+"}
	s.ViewSungJuk()
	s.InputSungJuck("영어", "A")
	s.ViewSungJuk()

}

func Increase(x *int) {
	*x++
}
