package animal

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Printf("%s: woof!\n", d.Name)
}
