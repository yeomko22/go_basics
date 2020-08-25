package ch3_struct

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type Player struct {
	Person
	Team string
}

func (p *Person) Introduce() {
	fmt.Print("Hello, this is my info: ")
	p.Info()
}

func (p *Person) Info() {
	fmt.Printf("name: %s, age: %d\n", p.Name, p.Age)
}

func (p *Player) Info() {
	fmt.Printf("name: %s, age: %d, team: %s\n", p.Name, p.Age, p.Team)
}

func SayHello(p Person) {
	fmt.Printf("Hello %s\n", p.Name)
}

func TestStruct(t *testing.T) {
	person := Person{"John", 30}
	player := Player{Person{"Lilard", 25}, "portland"}
	person.Info()
	player.Info()
	person.Introduce()
	player.Introduce()
	SayHello(person)
	SayHello(player)
}
