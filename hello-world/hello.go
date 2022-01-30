package main

import (
	"fmt"
	"runtime"
)

var rn rune = 88

const (
	LIMBO = 121
	GIMBO = "Basic"
)

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var xin int = 1

type animal struct {
	category string
	kingdom  string
}

type harbivorous struct {
	anim   animal
	eating string
}

type human interface {
	speak()
}

func humanCall(h human) {
	h.speak()
}
func main() {

	fmt.Printf("constant limbo %v and gimbo is %v ", LIMBO, GIMBO)
	n, _ := fmt.Println("Hello Go", "44", "true")
	fmt.Println(n)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%T\n", rn)

	str := "Hello, string"
	fmt.Printf("%T\n", str)

	bs := []byte(str)
	fmt.Println(bs)
	fmt.Printf("Type of string to byte conversion %T\n", bs)

	fmt.Printf("week starting sunday %d %d %d with a type %T\n", SUNDAY, MONDAY, TUESDAY, FRIDAY)

	for i := 0; i < 5; i++ {
		fmt.Println("this comes form for loop = ", i)
	}

	for xin < 5 {
		fmt.Println(xin)
		xin++
	}

	fmt.Println("xin after iteration =", xin)
	if x := 2; x == 2 {
		fmt.Println("2 is equal to 2")
	}

	switch {
	case false:
		fmt.Println("this should not show..")
		fallthrough
	case 2 == 4:
		fmt.Println("this compare should not appear")
	case 3 == 3:
		fmt.Println("this 3 equal is equal ")
		fallthrough
	case 5 == 6:
		fmt.Println("5 eqials 5..")
		fallthrough
	default:
		fmt.Println("Default")
	}

	// var color string = "yellow"
	switch "Blue" {
	//switch color {
	case "Hola", "Green", "yellow":
		fmt.Println("Yellow should print")
		fallthrough
	case "Red", "Blue":
		fmt.Println("Red and Blue")
	default:
		fmt.Println("None selected")
	}

	var x [5]int
	x[2] = 33
	fmt.Println(len(x))

	primes := []int{1, 2, 3}
	fmt.Println(primes)
	primes = append(primes, 4, 5, 6)

	for _, val := range primes[3:] {
		fmt.Println(val)

	}

	odd := []int{7, 9, 11}

	primes = append(primes, odd...)
	fmt.Println(primes)

	eud := make([]int, 5, 10)
	fmt.Println(eud)
	fmt.Println(len(eud))
	fmt.Println(cap(eud))

	actors := []string{"Tom", "Stall", "John"}
	actress := []string{"Julia", "Kris", "Kate"}

	act := [][]string{actors, actress}
	fmt.Println(act)

	mp := map[string]int{
		"Johhny":      1,
		"Walker Walk": 2,
	}

	fmt.Println(mp)
	mp["Tarzan"] = 3

	if v, ok := mp["Walker Walk"]; ok {
		fmt.Println(v)
	}

	if mk, ok := mp["Walker Walk"]; ok {
		fmt.Printf("Deleting Value is %v =", mk)
		delete(mp, "Walker Walk")
	}

	for kmp, vmp := range mp {
		fmt.Printf("key is %v and value is %v ", kmp, vmp)
	}

	a1 := animal{
		category: "Mammals",
		kingdom:  "root",
	}
	fmt.Println(a1)

	hb := harbivorous{
		anim:   a1,
		eating: "grass",
	}

	fmt.Println(hb)
	hb.speak()

	fmt.Println("Anonymous struct===")

	anon := struct {
		firstname string
		lastname  string
	}{
		firstname: "Tix",
		lastname:  "Wix",
	}

	fmt.Println(anon)

	gj, tq := woo("rim", "jhim")

	fmt.Printf("func call returned %v and %v", gj, tq)

	fp := myvariadic(3, 4, 5, 6, 7, 8)
	fmt.Printf("sum of value by variadic function is =%v ", fp)

	slint := []int{5, 6, 7, 8, 9}
	slintsum := myvariadic(slint[2:]...)
	fmt.Println(slintsum)
	defer foo()
	bar()

	humanCall(hb)
}
func foo() {
	fmt.Println("FOO is call")
	func(x int) {
		fmt.Println("Anonymous func called===", x)
	}(42)

	fc := func(s string) {
		fmt.Println(s)
	}

	fc("=== deoo ==")

	fg := hoo()
	fgint := fg()
	fmt.Println(fgint)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3sum := sum(s3...)
	fmt.Println("sum of all numbers =", s3sum)
	s3evensum := even(sum, s3...)
	fmt.Println("sum of even numbers = ", s3evensum)
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

func bar() {
	fmt.Println("BAR is call")
}

func hoo() func() int {
	f := func() int {
		return 45
	}
	return f
}
func woo(f string, g string) (string, bool) {
	a := fmt.Sprint("hello there", f, `says "Hola!! `, g)
	return a, true
}

func myvariadic(x ...int) int {
	fmt.Printf("type of variadic array is %T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func (hb harbivorous) speak() {
	fmt.Printf("Eating = %v and anim %v", hb.eating, hb.anim.category)
}
