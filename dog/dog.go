package dog

import "github.com/Roma4286/Lab4-Go/animal"

type Dog struct {
	animal.Animal
}

func (d Dog) Speak() string {
	return "Woof"
}