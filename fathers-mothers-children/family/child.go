package family

import "fmt"

type Child struct {
	Name     string
	Age      int
	IsCalm   bool
	IsHungry bool
}

func NewChild(name string, age int) *Child {
	return &Child{Name: name, Age: age, IsCalm: true, IsHungry: false}
}

func (c *Child) CalmDown() {
	c.IsCalm = true
}

func (c *Child) GetHungry() {
	c.IsHungry = true
}

func (c *Child) Feed() {
	c.IsHungry = false
}

func (c *Child) Info() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Спокоен: %t, Голоден: %t", c.Name, c.Age, c.IsCalm, c.IsHungry)
}
