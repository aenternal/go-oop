package main

import (
	"fmt"
	"log"

	"go-oop/fathers-mothers-children/family" // замените на правильный путь
)

func main() {
	// Создаем родителя
	parent := family.NewParent("Анна", 25)
	fmt.Println(parent.Info())

	// Создаем детей
	child1 := family.NewChild("Миша", 9)
	child2 := family.NewChild("Оля", 5)

	// Добавляем детей к родителю
	if err := parent.AddChild(child1); err != nil {
		log.Fatalf("Ошибка добавления ребенка: %v", err)
	}
	if err := parent.AddChild(child2); err != nil {
		log.Fatalf("Ошибка добавления ребенка: %v", err)
	}

	// Выводим информацию о родителе и детях
	fmt.Println(parent.Info())
	fmt.Println(child1.Info())
	fmt.Println(child2.Info())

	child1.GetHungry()
	child1.IsCalm = false
	fmt.Printf("%s проголодался и орёт!\n", child1.Name)

	// Родитель успокаивает и кормит ребенка
	parent.CalmChild(child1)
	parent.FeedChild(child1)

	// Выводим обновленную информацию о ребенке
	fmt.Println(child1.Info())
}
