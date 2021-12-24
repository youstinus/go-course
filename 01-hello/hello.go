package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	foo()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	declare()
	zeroValues()
	types()
	conversion()

	bar()
}

func foo() {
	fmt.Println("Im foo")
}

func bar() {
	fmt.Println("Exited")
}

// works
var z = 5
var k int
var s = ` "asdasd" asda sd`

// does not work here:
// c := 5

func declare() {
	x := 42
	fmt.Println(x)
	x = 80
	fmt.Println(x)
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)
	k = 6
	fmt.Println(k)
	fmt.Printf("%T\n", z)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

func zeroValues() {
	var c int
	fmt.Println(c)
	var c1 string
	fmt.Println(c1)

}

func types() {
	var a int
	type hotdog int
	var b hotdog
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

func conversion() {
	var a int
	type hotdog int
	var b hotdog
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	a = int(b) // conversion
	fmt.Printf("%T\n", a)
}
