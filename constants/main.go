package main

import "fmt"

const (
	DATE_FORMAT = "mm/dd/yyyy"
	localConst  = 34
) //the varaible which starts with smallcase have only localscope
var (
	t = 23
	i = 1 //this is define in global scope and undefined
) //theyhaveglobalscopewecanaccessitanywhereinthwfike

func main() {
	const name string = "Rajesh"
	const name1 = "sri"
	const i = 200
	t = 56 //updating the t,which is defined in global variable
	fmt.Println(name)
	fmt.Println(DATE_FORMAT, localConst, i)
}
