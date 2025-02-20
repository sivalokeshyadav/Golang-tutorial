package main

import "fmt"

func main() {
	//create array
	var ai [10]int
	for _, v := range ai {
		fmt.Printf("%v,", v)
	}
	ai[0] = 1
	ai[4] = 5
	ai[9] = 10
	//printArray(ai)
	// bi := [10]int{}
	// for i, v := range bi {

	// 	fmt.Printf("%v:%v,", i, v)
	// }
	//arrWithValues := [5]int{1, 2, 3, 4, 5}
	arrWithValues := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrWithValues)
	//simpleslicearray
	slice1 := ai[0:4]
	fmt.Println("Length:", len(slice1), "Contents:", slice1, "capacity:", cap(slice1))
	//create slice using make
	s1 := make([]int, 10)
	fmt.Println("Length:", len(s1), "Contents:", s1, "capacity:", cap(s1))
	s1 = append(s1, 12)
	fmt.Println("Length:", len(s1), "Contents:", s1, "capacity:", cap(s1))
	//create a slice
	slice2 := []int{}
	fmt.Println("Length:", len(slice2), "Contents:", slice2, "capacity:", cap(slice2))
	//copy slice
	slice3 := append(slice2, 12)
	fmt.Println("Length:", len(slice3), "Contents:", slice3, "capacity:", cap(slice3))
	sArray := [10]string{"array"}
	fmt.Println((sArray))
	updateArray(sArray)
	fmt.Println(sArray)
	slice4 := sArray[:]
	updateSlice(slice4)
}

//arraysarevaluetyped-->copiedtothefunction
func updateArray(as [10]string) {
	as[0] = "Updated"
}
func updateSlice(s1 []string) {
	s1[0] = "Updated"
}
func printArray(ai [10]int) {
	for i, v := range ai {
		fmt.Printf("%v:%v,", i, v)
	}

}
