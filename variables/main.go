package main

import "fmt"

func main() {
	//var is a keyword
	//foo is the name of the variable
	//intishebuiltintype
	//42 is the value we are initialzing with

	var foo int = 42    //declaration with initliazation
	var bar int         //declaration without initalization it takes value as 0
	var i, j int = 2, 1 //muitlple variables declaration with initliazation
	var x, y int        //muitlple variables declaration without initliazation
	name := "Rajesh"    //var name string="Rajesh"
	a, b := 2, "Siva"   //go is a strictly types lang,we cant declare varialbe again when we intialze it
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(i, j)
	fmt.Println(x, y)
	fmt.Println(name)
	fmt.Println(a, b)
}
