package main

import "fmt"

func main() {
	var flag bool = true
	var i int = -123 //int,int8,int16,int32,int64
	var str string = "Some string"
	var ui uint = 123   //uint,uint8,uint16,uint32,uint64
	var r = 65346       //rune=int32
	var b byte = 255    //byte==uint8
	var f float32 = 3.5 //float32,float64
	var c int8 = -127   //-127 to 127
	fmt.Println(flag)
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(ui)
	fmt.Println(b)
	fmt.Println(r)
	fmt.Println(f)
	fmt.Println(c)
}
