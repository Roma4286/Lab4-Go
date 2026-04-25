package animal

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Some sound"
}