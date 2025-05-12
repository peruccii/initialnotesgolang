package main

import (
	"encoding/json"
	"fmt"
)

type Type struct {
	Color string `json:"color"`
}

type Foo struct {
	Name string `json:"name"`
	TypeFoo Type `json:"typeFoo"`
}

type rect struct {
	x, y int
}

func (r *rect) area() int {
	return r.x * r.y
}

// function]
// func (type) nameFunc {}
func (f Foo) Show() {
	fmt.Printf("Name: %s", f.Name)
}

type Boat struct {
	Name string
	occupants []string
}

func (b *Boat) AddOccupant(name string) *Boat {
	b.occupants = append(b.occupants, name)
	return b
}

func (b *Boat) RemoveOccupant(name string) {
	for index, occupant := range b.occupants {
		if occupant == name {
			b.occupants = append(b.occupants[:1],  b.occupants[index+1:]... )
			return 
		}
	}
}

func (b Boat) Manifest() {
	fmt.Println("The", b.Name, "has the following occupants:")
	for index, item := range b.occupants {
		fmt.Println("\t", index, item)
	}
}

type Submersible interface {
	Dive()
}

type Shark struct {
	Name string
	isUnderWater bool
}

func (s *Shark) Dive() {
	s.isUnderWater = true
}

func (s Shark) String() string {
	if s.isUnderWater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

func submerge(s Submersible) {
	s.Dive()
}

func main() {
	
	s := &Shark{
		Name: "Nice Shark",
	}

	submerge(s)
	fmt.Println(s)
	b := &Boat{
		Name: "S.sDigitalOcena",
	}

	b.AddOccupant("Im the first occupannt")
	b.AddOccupant("Im the second occupannt")
	b.RemoveOccupant("Im the second occupannt")
	b.Manifest()
	
	f := Foo{ // creating 
		Name: "Fooling",
		TypeFoo: Type{
			Color: "some color here",
		},
	}

	fJson, err := json.Marshal(f) // struct to JSON
	
	if err != nil {
		panic("error")
	}

	fmt.Println(string(fJson))

	r := rect{x: 20, y: 2} // creating
	
	fmt.Println("Area", r.area())
}