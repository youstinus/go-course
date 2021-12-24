package main

import (
	"fmt"
)

func main() {
	assign()
	declare()
	assign2()
	ownType()
	fourth()
	fifth()
}

func assign() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
}

var x int
var y string
var z bool

func declare() {
	fmt.Println(x, y, z)
}

func assign2() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

type a int

var xx a

func ownType() {
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)
	xx = 5
	fmt.Println(xx)
}

var yy int

func fourth() {
	fmt.Println(yy)
	fmt.Printf("%T\n", yy)
	yy = int(xx)
	fmt.Println(yy)

}

func fifth() {

}
