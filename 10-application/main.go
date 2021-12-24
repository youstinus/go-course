package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

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

type person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func func0() {
	p1 := person{
		First: "James",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Age:   122,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

	var ob []person
	err = json.Unmarshal(bs, &ob)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ob)

}

func func1() {

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func func2() {
	p1 := person{"James", 32}
	p2 := person{"Money penny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

func func3() {
	pass := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

	given := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(given))

	if err != nil{
		fmt.Println("Cant login", err)
		return
	}

	fmt.Println("You're logged in")
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
