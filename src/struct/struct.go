package main

import "fmt"

type Person struct { // Person이라는 type을 만든 것
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjuk)
}

func ViewSungJuk(s Student) {
	fmt.Println(s.sungjuk)
}

func (s Student) InputSungJuck(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func main() {
	var p Person // Person이라는 type을 만든 것
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 29}
	p3 := Person{name: "Jhon"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 24

	p1.PrintName()

	var a Student
	a.name = "철수"
	a.class = 1
	a.sungjuk.name = "수학"
	a.sungjuk.grade = "A"

	a.ViewSungJuk()
	ViewSungJuk(a)

	a.InputSungJuck("과학", "B") // 복사되어 들어가므로 변경 x / pointer 필요!
	a.ViewSungJuk()

}
