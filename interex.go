package main

import (
	"fmt"

	"github.com/pdk/interex/animal"
	"github.com/pdk/interex/ghost"
)

type speaker interface {
	Speak()
}

type floater interface {
	Name() string
	Float()
}

func main() {

	speakers := []speaker{}

	speakers = append(speakers, animal.Dog{Name: "Spot"})
	speakers = append(speakers, ghost.New("Abe", "Lincoln"))
	speakers = append(speakers, animal.Cat{Name: "Miffy"})

	for _, s := range speakers {
		s.Speak()
	}

	for _, s := range speakers {
		switch v := s.(type) {
		case floater:
			fmt.Printf("%s is a floater!\n", v.Name())
		}
	}
}
