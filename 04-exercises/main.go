package main

import "fmt"

func main() {
	first()
	second()
	third()
	literal()
	iotas()
}

func first() {
	n := 6516
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}

func second() {
	// useless
}

func third() {

}

func literal() {
	a := `asdasda
			asdasda "asd'asd 'asd w454678iuytrb


sdfsdf			 	eq wd"`
	fmt.Println(a)
}

func iotas() {
	const (
		a = 2014 + iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
