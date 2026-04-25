package cat

import "example.com/animal"

type Cat struct {
	animal.Animal
}

func (c Cat) Speak() string {
	return "Meow"
}