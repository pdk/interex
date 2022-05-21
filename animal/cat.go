package animal

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("%s: meow!\n", c.Name)
}
