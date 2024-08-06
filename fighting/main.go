package main

import (
	"go-oop/fighting/battle"
	"go-oop/fighting/warrior"
)

func main() {
	redWarrior := warrior.NewWarrior("Красный")
	blueWarrior := warrior.NewWarrior("Синий")

	battle.StartBattle(redWarrior, blueWarrior)
}
