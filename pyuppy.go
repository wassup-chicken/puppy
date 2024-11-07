package puppy

import (
	"fmt"

	"github.com/wassup-chicken/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}

func From12() {
	fmt.Println("I'm from verision 1.2.0")
}
func From13() {
	fmt.Println("I'm from verision 1.3.0")
}
func From14() {
	fmt.Println("I'm from verision 1.4.0")
}
