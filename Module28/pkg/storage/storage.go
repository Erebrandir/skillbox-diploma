package storage

import (
	"awesomeProject/Module28/pkg/student"
	"errors"
)

type StudentMap map[string]*student.Student

func NewStorage() StudentMap {
	return make(map[string]*student.Student)
}

func (s StudentMap) Put(student *student.Student) {
	s[student.Name()] = student
}

func (s StudentMap) Get(studentName string) (*student.Student, error) {
	student, ok := s[studentName]
	if !ok {
		return nil, errors.New("Студент в базе не найден!")
	} else {
		return student, nil
	}
}
