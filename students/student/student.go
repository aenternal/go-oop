package student

import (
	"fmt"
	"sort"
)

type Student struct {
	FullName    string
	GroupNumber string
	Grades      []int
}

func NewStudent(fullName string, groupNumber string, grades []int) (*Student, error) {
	if len(grades) != 5 {
		return nil, fmt.Errorf("успеваемость должна содержать ровно 5 элементов")
	}
	for _, grade := range grades {
		if grade < 2 || grade > 5 {
			return nil, fmt.Errorf("оценки должны быть в диапазоне от 2 до 5")
		}
	}
	return &Student{FullName: fullName, GroupNumber: groupNumber, Grades: grades}, nil
}

func (s *Student) AverageGrade() float64 {
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func SortByAverageGrade(students []*Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].AverageGrade() < students[j].AverageGrade()
	})
}
