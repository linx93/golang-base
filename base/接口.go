package main

import "fmt"

// USB defined a USB interface,it can read and write
type USB interface {
	read()
	write()
}

type Computer struct {
	Name string
}

func (computer Computer) read() {
	fmt.Println(computer.Name, " read ")
}
func (mobile Mobile) write() {
	fmt.Println(mobile.Name, " write ")
}

type Mobile struct {
	Name string
}

/*func main() {
	computer := Computer{
		Name: "Apple conputer ",
	}
	computer.read()

	mobile := Mobile{
		Name: "huawei",
	}
	mobile.write()

}*/
