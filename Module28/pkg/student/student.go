package student

import "strconv"

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) Name() string {
	return s.name
}

func NewStudent(name string, age int, grade int) *Student {
	return &Student{
		name:  name,
		age:   age,
		grade: grade,
	}
}

func (s *Student) Info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}
