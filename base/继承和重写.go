package main

import "fmt"

type animal struct {
}

func (a *animal) Name() {
	fmt.Println("this is animal")
}

type dog struct {
	animal //Dog 构造体继承 Animal
}

//重写Animal的Name方法
//func (a *dog) Name() {
//	fmt.Println("this is dog")
//}

/*func main() {
	dog := dog{}
	dog.animal.Name() // 也可以点出父构造体
	dog.Name()        // 可以直接点出.Name()

	//pring result

	//this is animal
	//this is animal

	//重写Animal的Name方法

	//this is animal
	//this is dog

}*/
