package main

import (
	"fmt"
)

func main() {
	arrays()
	slice1()
	slice2()
	slice3()
	slice4()
	slice5()
	slice6()
	map1()
}

func arrays() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

}

func slice1() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}

func slice2() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println("slice", x)
	fmt.Println("length", len(x))
	fmt.Println("first", x[0])

	for i, v := range x {
		fmt.Println(i, v)
	}
}

func slice3() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println("slice", x)
	fmt.Println("from 1(include) to 3(exclude)", x[1:3])
	fmt.Println("from 1(include) to end", x[1:])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i]) // same as range
	}
}

// append
func slice4() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println("slice", x)
	x = append(x, 9, 651, 5165, 65, 1)
	fmt.Println("appended", x)
	y := []int{5, 46, 4, 8}
	x = append(x, y...)
	fmt.Println("appended", x)
	x = append(x[:5], x[8:]...)
	fmt.Println("appended", x)
}

func slice5() {

	x := make([]int, 3, 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func slice6() {
	x := []string{"JAmes"}
	fmt.Println(x)
	y := []string{"asdasd"}
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)

}

func map1() {
	m := map[string]int{
		"James": 32,
		"MISS":  27,
	}
	fmt.Println(m)
	fmt.Println(m["James"], m["what"])

	v, ok := m["Bas"]
	fmt.Println(v, ok)

	m["todd"] = 33

	if v, ok := m["asda"]; ok {
		fmt.Println("Good", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "todd")
	delete(m, "toddsasdasd")

	fmt.Println(m)

}
