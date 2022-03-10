package storage

import (
	"awesomeProject/Module28/pkg/student"
	"errors"
)

type Storage map[string]*student.Student

func NewStorage() Storage {
	return make(map[string]*student.Student)
}

func (s Storage) Put(student *student.Student) {
	s[student.Name()] = student
}

func (s Storage) Get(studentName string) (*student.Student, error) {
	student, ok := s[studentName]
	if !ok {
		return nil, errors.New("Студент не найден!")
	} else {
		return student, nil
	}
}
