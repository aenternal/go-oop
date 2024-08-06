package warrior

import "fmt"

type Warrior struct {
	Health int
	damage int
	Name   string
}

func NewWarrior(name string) *Warrior {
	return &Warrior{Health: 100, damage: 20, Name: name}
}

func (w *Warrior) Attack(target *Warrior) {
	target.Health -= w.damage
	fmt.Printf("%s атакует %s и наносит %d урона. У %s осталось %d здоровья.\n", w.Name, target.Name, w.damage, target.Name, target.Health)
}
