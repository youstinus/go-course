package main

import "fmt"

func main() {
	loops()
	stringing()
	switching()
}

func loops() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(x)
	}

	for {
		fmt.Println("while loop")
		break
	}

	c := 0
	for {
		c++
		fmt.Println("while loop")
		if c <= 1 {
			continue
		} else if c > 1 {
			break
		}
	}
}

func stringing() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v\t%#x\t%#U\t%s\n", i, i, i, string(i))
	}
}

func switching() {
	switch {
	case false:
		fmt.Println("la")
		//break
	case 2 != 4:
		fmt.Println("asdasd")
		//break
	default:
		fmt.Println("qwertghx")
		break
	}
}
