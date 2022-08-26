package main

// NewInt 类型定义
type NewInt int

// MyInt 类型别名
type MyInt = int

//
//func main() {
//	//类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别
//	var a NewInt
//	var b MyInt
//	fmt.Printf("a type=%T\n", a)
//	fmt.Printf("b type=%T\n", b)
//	//输出类型
//	//a type=main.NewInt
//	//b type=int
//
//	//结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
//
//	//四、总结
//	//类型定义:type 类型名 类型 type NewInt int, 自定义类型和原类型是两种不同的类型,会创建新类型
//
//	//类型别名:type 类型名 = 类型 type MyInt = int,类型别名和原类型一样,并没有创建新类型
//	//在当下的阶段，必将由程序员来主导，甚至比以往更甚。
//}
