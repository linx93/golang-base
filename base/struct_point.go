package main

type Person struct {
	Age  int
	Name string
}

//
//func main() {
//	person := Person{
//		Age:  -1,
//		Name: "linx",
//	}
//	fmt.Println(person)
//	SetAge(person)
//	fmt.Println(person)
//	SetName(&person)
//	fmt.Println(person)
//
//	//print result
//	//{-1 linx}
//	//{-1 linx}
//	//{-1 tom}
//
//	person.EatFood("apple")
//	person.EatFood1("apple")
//}
//
//func SetAge(person Person) {
//	person.Age = 100
//}
//func SetName(person *Person) {
//	person.Name = "tom"
//}
//
//func (person Person) EatFood(foodName string) {
//	fmt.Println(person.Name, " eat ", foodName)
//}
//func (person *Person) EatFood1(foodName string) {
//	fmt.Println(person.Name, " eat ", foodName)
//}
