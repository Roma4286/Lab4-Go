package dog

import "example.com/animal"

type Dog struct {
	animal.Animal
}

func (d Dog) Speak() string {
	return "Woof"
}