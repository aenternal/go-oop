package family

import "fmt"

type Parent struct {
	Name     string
	Age      int
	Children []*Child
}

func NewParent(name string, age int) *Parent {
	return &Parent{Name: name, Age: age, Children: []*Child{}}
}

func (p *Parent) AddChild(child *Child) error {
	if child.Age > p.Age-16 {
		return fmt.Errorf("возраст ребенка должен быть меньше возраста родителя хотя бы на 16 лет")
	}
	p.Children = append(p.Children, child)
	return nil
}

func (p *Parent) Info() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Количество детей: %d", p.Name, p.Age, len(p.Children))
}

func (p *Parent) CalmChild(child *Child) {
	child.CalmDown()
	fmt.Printf("%s успокоил(а) %s\n", p.Name, child.Name)
}

func (p *Parent) FeedChild(child *Child) {
	child.Feed()
	fmt.Printf("%s покормил(а) %s\n", p.Name, child.Name)
}
