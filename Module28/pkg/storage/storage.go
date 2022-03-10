package storage

import (
	"awesomeProject/Module28/pkg/student"
	"errors"
)

type University struct {
	studentMap map[string]*student.Student
}

func NewUniversity() *University {
	return &University{
		studentMap: make(map[string]*student.Student),
	}
}

func (u *University) Put(student *student.Student) error {
	if _, found := u.studentMap[student.Name()]; !found {
		u.studentMap[student.Name()] = student
		return nil
	}
	return errors.New("Данный студент уже есть в списке!")
}

func (u *University) Get() []*student.Student {
	var students []*student.Student
	for _, v := range u.studentMap {
		students = append(students, v)
	}
	return students
}
