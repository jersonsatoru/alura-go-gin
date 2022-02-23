package repositories

import (
	"github.com/jersonsatoru/alura-go-gin/internal/db"
	"github.com/jersonsatoru/alura-go-gin/internal/entities"
	"github.com/jersonsatoru/alura-go-gin/internal/orm"
)

type StudentRepository struct{}

type IStudentRepository interface {
	Create(*entities.Student) (*entities.Student, error)
	Find() []entities.Student
}

func (s *StudentRepository) Create(student *entities.Student) (*entities.Student, error) {
	studentORM := orm.Student{
		Name: student.Name,
		CPF:  student.CPF,
		RG:   student.RG,
	}
	tx := db.DB.Create(&studentORM)
	if err := tx.Error; err != nil {
		return nil, err
	}
	student.ID = int64(studentORM.ID)
	return student, nil
}

func (s *StudentRepository) Find() []entities.Student {
	var studentsORM []orm.Student
	db.DB.Find(&studentsORM)
	var students []entities.Student
	for _, s := range studentsORM {
		students = append(students, entities.Student{
			ID:   int64(s.ID),
			Name: s.Name,
			CPF:  s.CPF,
			RG:   s.RG,
		})
	}
	return students
}
