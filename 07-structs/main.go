package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	do2()
	do3()
}

func do2() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "James2",
		last:  "Bond",
		age:   62,
	}
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "James2",
		last:  "Bond",
		age:   62,
	}
	sa1 := secretAgent{
		person: person{
			first: "Sup",
			last:  "last",
			age:   4,
		},
		first: "yo",
		ltk:   true,
	}
	fmt.Println(p1, p2, p3, sa1)
	fmt.Println(sa1.person.first, sa1.last, sa1.age)
}

func do3() {

}
