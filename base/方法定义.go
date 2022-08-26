package main

import (
	"fmt"
	"strconv"
)

type Animal struct {
	Name string
}

/*
func main() {
	dog := Animal{Name: "dog"}
	dog.run(1)
	fmt.Println(dog)

	fmt.Println("-----------------------------")

	birdPoint := &Animal{Name: "bird"}
	birdPoint.fly(1)
	fmt.Println(*birdPoint)

	// print result
	//dog run 1 hours
	//{dog}
	//-----------------------------
	//bird fly 1 hours
	//{changed bird}
}*/
func (animal Animal) run(hours int) string {
	run := animal.Name + " run " + strconv.Itoa(hours) + " hours"
	fmt.Println(run)
	//change animal
	animal.Name = "changed " + animal.Name
	return run
}

func (animal *Animal) fly(hours int) string {

	fly := animal.Name + " fly " + strconv.Itoa(hours) + " hours"
	fmt.Println(fly)
	//change animal
	animal.Name = "changed " + animal.Name
	return fly
}
