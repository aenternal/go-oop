package main

import (
	"fmt"
	"log"

	"go-oop/students/student"
)

func main() {
	var students []*student.Student

	data := []struct {
		fullName    string
		groupNumber string
		grades      []int
	}{
		{"Иванов Иван Иванович", "1", []int{5, 4, 3, 4, 5}},
		{"Петров Петр Петрович", "2", []int{3, 2, 4, 3, 3}},
		{"Сидоров Сидор Сидорович", "3", []int{5, 5, 5, 4, 5}},
		{"Кузнецова Ольга Петровна", "4", []int{4, 4, 4, 4, 4}},
		{"Смирнова Анна Викторовна", "5", []int{2, 3, 2, 3, 2}},
		{"Морозов Сергей Николаевич", "6", []int{5, 4, 3, 5, 4}},
		{"Новикова Елена Ивановна", "7", []int{3, 3, 4, 3, 3}},
		{"Волков Дмитрий Павлович", "8", []int{5, 5, 4, 4, 5}},
		{"Зайцева Мария Александровна", "9", []int{4, 3, 4, 4, 3}},
		{"Попов Александр Сергеевич", "10", []int{5, 4, 5, 5, 5}},
	}

	for _, d := range data {
		createdStudent, err := student.NewStudent(d.fullName, d.groupNumber, d.grades)
		if err != nil {
			log.Fatalf("Ошибка при создании студента %s: %v", d.fullName, err)
		}
		students = append(students, createdStudent)
	}

	student.SortByAverageGrade(students)

	fmt.Println("Отсортированный список студентов по среднему баллу:")
	for _, s := range students {
		fmt.Printf("ФИ: %s, Группа: %s, Средний балл: %.2f\n", s.FullName, s.GroupNumber, s.AverageGrade())
	}
}
