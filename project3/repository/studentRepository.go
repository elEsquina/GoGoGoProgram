package repository

import (
	"errors"
	"project3/model"
	"fmt"
)

type StudentRepository struct {
	Students []model.Student
}

func NewStudentRepository() *StudentRepository {
	LoadJson()
	fmt.Println("Loading")
	fmt.Println(ServerStorage)
	return &StudentRepository{
		Students: ServerStorage,
	}
}

func (s *StudentRepository) AddStudent(student model.Student) error {
	for _, existingStudent := range s.Students {
		if existingStudent.Id == student.Id {
			return errors.New("ID already exists")
		}
	}
	s.Students = append(s.Students, student)

	return SaveJson(s.Students) 
}

func (s *StudentRepository) GetStudentById(id int) (model.Student, error) {
	for _, student := range s.Students {
		if student.Id == id {
			return student, nil
		}
	}
	return model.Student{}, errors.New("student not found")
}

func (s *StudentRepository) UpdateStudent(student model.Student, id int) error {
	for i, existingStudent := range s.Students {
		if existingStudent.Id == id {
			s.Students[i] = student
			return SaveJson(s.Students)
		}
	}
	return errors.New("student not found")
}

func (s *StudentRepository) DeleteStudent(id int) error {
	for i, student := range s.Students {
		if student.Id == id {
			s.Students = append(s.Students[:i], s.Students[i+1:]...)
			return SaveJson(s.Students)
		}
	}
	return errors.New("student not found")
}

func (s *StudentRepository) GetStudents() ([]model.Student) {
	return s.Students
}
