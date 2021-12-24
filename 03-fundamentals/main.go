package main

import (
	"fmt"
	"runtime"
)

func main() {
	variables()
	strings()
	constants()
	iotas()
	shifting()
	shifting2()
	shifting3()
}

var x bool

func variables() {
	r := "ąčęėįšųūž"
	fmt.Println(r, x)
	fmt.Println(runtime.GOARCH)
}

func strings() {
	s := "Hello, ąčęėįšųū9ž"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []rune(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}

const a = 42
const (
	b = 42.78
)
const (
	c string = "James Bond"
)

func constants() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}

const (
	d = iota
	e
	f
)

func iotas() {
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}

func shifting() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

}

func shifting2() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)

}

func shifting3() {
	const (
		_  = iota             // 0
		kb = 1 << (iota * 10) // 1
		mb = 1 << (iota * 10) // 2
		gb = 1 << (iota * 10) // 3
	)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)

}
