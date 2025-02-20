package main

import "fmt"

func main() {
	//calling the function with {}
	simpleFunction()
	f := simpleFunction
	f()
	withParams("siva", 2)
	withParamsWithSameType(5, 2)
	age := returningSingleValue()
	fmt.Println("Age:", age)
	a, b := returningzMultipleValue()
	fmt.Println(a, b)
	a, b = returningzMultipleValue()
	fmt.Println(a, b)
}

//defininga simple function with no return types
func simpleFunction() {
	fmt.Println("Simple function")
}
func withParams(p1 string, p2 int) {
	fmt.Println("Inside with params=:", p1, p2)
}
func withParamsWithSameType(p1, p2 int) {
	fmt.Println("Inside with params with sametype=:", p1, p2)
}
func returningSingleValue() int {
	return 43
}
func returningzMultipleValue() (int, string) {
	return 43, "siva"
}
func returningzMultipleValue2() (n int, s string) {
	s = "siva"
	n = 23
	return
}
