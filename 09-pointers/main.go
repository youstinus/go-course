package main

import "fmt"

func main() {
	func0()
	func1()
	func2()
	func3()
	func4()
	func5()
	func6()
	func7()
	func8()
	func9()
}

func func0() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}

func func1() {
	x := 0
	foo(x)
	fmt.Println(x)

	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	fooPointer(&x)
	fmt.Println("x  after", &x)
	fmt.Println("x  after", x)

}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func fooPointer(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y  after", y)
	fmt.Println("y  after", *y)
}

func func2() {

}

func func3() {

}

func func4() {

}

func func5() {

}

func func6() {

}

func func7() {

}

func func8() {

}

func func9() {

}
