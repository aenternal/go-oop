package battle

import (
	"fmt"
	"go-oop/fighting/warrior"
	"math/rand"
)

func StartBattle(redWarrior, blueWarrior *warrior.Warrior) {
	attacks := []func(*warrior.Warrior){
		redWarrior.Attack,
		blueWarrior.Attack,
	}

	for redWarrior.Health > 0 && blueWarrior.Health > 0 {
		attackerIndex := rand.Intn(len(attacks))
		var attacker, target *warrior.Warrior

		if attackerIndex == 0 {
			attacker = redWarrior
			target = blueWarrior
		} else {
			attacker = blueWarrior
			target = redWarrior
		}

		attacker.Attack(target)

		if redWarrior.Health <= 0 {
			fmt.Printf("%s проиграл!", redWarrior.Name)
			break
		}
		if blueWarrior.Health <= 0 {
			fmt.Printf("%s проиграл!", blueWarrior.Name)
			break
		}
	}
}
