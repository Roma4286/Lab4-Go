package cat

import "github.com/Roma4286/Lab4-Go/animal"

type Cat struct {
	animal.Animal
}

func (c Cat) Speak() string {
	return "Meow"
}