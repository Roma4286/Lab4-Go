package main

import (
	"fmt"
	"example.com/dog"
	"example.com/cat"
)

func main() {
	d := dog.Dog{}
	c := cat.Cat{}

	fmt.Println(d.Speak())
	fmt.Println(c.Speak())
}