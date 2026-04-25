package main

import (
	"fmt"
	"github.com/Roma4286/Lab4-Go/dog"
	"github.com/Roma4286/Lab4-Go/cat"
)

func main() {
	d := dog.Dog{}
	c := cat.Cat{}

	fmt.Println(d.Speak())
	fmt.Println(c.Speak())
}