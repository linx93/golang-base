package main

import "fmt"

func main() {
	str := "i am linx"
	var strPointType *string
	fmt.Println("-------------------------------------")
	fmt.Printf("str_point_type: %v\n", strPointType)
	fmt.Println("-------------------------------------")
	strPointType = &str
	fmt.Printf("str_point_type: %T\n", strPointType)
	fmt.Printf("str_point_type: %v\n", strPointType)
	fmt.Printf("str_point_type: %v\n", *strPointType)
	fmt.Println("-------------------------------------")
	str = "i am hwm"
	fmt.Printf("str_point_type: %T\n", strPointType)
	fmt.Printf("str_point_type: %v\n", strPointType)
	fmt.Printf("str_point_type: %v\n", *strPointType)
	fmt.Println("-------------------------------------")
	str1 := "hello world"
	strPointType = &str1
	fmt.Printf("str_point_type: %T\n", strPointType)
	fmt.Printf("str_point_type: %v\n", strPointType)
	fmt.Printf("str_point_type: %v\n", *strPointType)
}
