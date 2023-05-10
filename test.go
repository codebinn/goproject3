package main

import "fmt"

type Aaa struct {
	num int
}

func (a Aaa) afunc() {
	a.num = 33
	fmt.Println("afunc:", a.num)
}

func (a *Aaa) bfunc() {
	a.num = 10
	fmt.Println("bfunc:", a.num)
}

func main() {
	var b Aaa
	a := &b
	fmt.Println("111", a.num)
	a.afunc()
	fmt.Println("333", a.num)
	a.bfunc()
	fmt.Println("555", a.num)
}
