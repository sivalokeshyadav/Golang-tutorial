package main

import "fmt"

func main() {
	printSomething := func() {
		fmt.Println("something")
	}
	printSomething()
	//add := func() {

	//}
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%v+%v=%v", 2, 3, add(2, 3))
	func() {
		fmt.Println(("print siva something"))
	}()

}

// //above printfunction and below are same
// func printSomething() {
// 	fmt.Println("something1")
// }
