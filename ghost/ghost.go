package ghost

import "fmt"

type Ghost struct {
	firstName string
	lastName  string
}

func New(f, l string) Ghost {
	return Ghost{
		firstName: f,
		lastName:  l,
	}
}

func (g Ghost) Name() string {
	return fmt.Sprintf("%s %s", g.firstName, g.lastName)
}

func (g Ghost) Speak() {
	fmt.Printf("%s %s: boo!\n", g.firstName, g.lastName)
}

func (g Ghost) Float() {
	fmt.Printf("%s %s is floating ...\n", g.firstName, g.lastName)
}
