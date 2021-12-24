package main

import (
	"fmt"
)

func main() {
	func0(2, 3, 4, 5, 6, 7, 89, 8)
	defer func2()
	func1()
	func3()
	func5()
	func6()
	func7()
	func8()
	func9()
}

func func0(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	return x[0]
}

func func1() {
	xi := []int{2, 3, 4, 54, 68, 684, 86, 4}
	x := func0(xi...)
	fmt.Println(x)
}

func func2() {
	fmt.Println("Hello defer2")
}

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

type hotdog int

func func3() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(sa1)
	sa1.func4()
	p1.func4()
	bar(sa1)
	bar(p1)

	var x hotdog = 42
	fmt.Println(x)
	var y = int(x)
	fmt.Println(y)
}

func (s secretAgent) func4() {
	fmt.Println("im", s.first, s.last)
}
func (p person) func4() {
	fmt.Println("im", p.first, p.last)
}

type human interface {
	func4()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("asserted to person", h.(person).first)
	case secretAgent:
		fmt.Println("asserted to secret agent", h.(secretAgent).first)
	}

	fmt.Println("Im human", h)

}

func func5() {

	func() {
		fmt.Println("Anonymous func")
	}()

	func(x int) {
		fmt.Println("Anonymous func with param", x)
	}(42)

	f := func() {
		fmt.Println("Func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("Func expression with param", x)
	}
	f2(45)
}

func func6() {
	s1 := getString()
	fmt.Println(s1)

	s2 := func() int {
		return 456
	}()
	fmt.Println(s2)

	s3 := bar2()
	s4 := s3()
	fmt.Println(s4)
}

func getString() string {
	s := "Hello world"
	return s
}

func bar2() func() int {
	return func() int {
		return 464
	}
}

func func7() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("All numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("even numbers", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {

		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func func8() {
	a := counter()
	b := counter()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func counter() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	result := 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}

func func9() {
	n := factorial(4)
	fmt.Println(n)
	k := factorialLoop(4)
	fmt.Println(k)
}
