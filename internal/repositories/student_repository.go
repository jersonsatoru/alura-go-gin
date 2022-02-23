package repositories

import (
	"context"
	"time"

	"github.com/jersonsatoru/alura-go-gin/internal/db"
	"github.com/jersonsatoru/alura-go-gin/internal/entities"
	"github.com/jersonsatoru/alura-go-gin/internal/orm"
	"gorm.io/gorm"
)

type StudentRepository struct{}

type IStudentRepository interface {
	Create(*entities.Student) (*entities.Student, error)
	Find() ([]entities.Student, error)
	FindByID(int64) (*entities.Student, error)
	Delete(int64) error
	Update(*entities.Student) (*entities.Student, error)
}

func (s *StudentRepository) Create(student *entities.Student) (*entities.Student, error) {
	studentORM := orm.Student{
		Name: student.Name,
		CPF:  student.CPF,
		RG:   student.RG,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tx := db.DB.WithContext(ctx).Create(&studentORM)
	if err := tx.Error; err != nil {
		return nil, err
	}
	student.ID = int64(studentORM.ID)
	return student, nil
}

func (s *StudentRepository) Find() ([]entities.Student, error) {
	var studentsORM []orm.Student
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tx := db.DB.WithContext(ctx).Find(&studentsORM)
	if err := tx.Error; err != nil {
		return nil, err
	}
	var students []entities.Student
	for _, s := range studentsORM {
		students = append(students, entities.Student{
			ID:   int64(s.ID),
			Name: s.Name,
			CPF:  s.CPF,
			RG:   s.RG,
		})
	}
	return students, nil
}

func (s *StudentRepository) FindByID(id int64) (*entities.Student, error) {
	var studentORM orm.Student
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tx := db.DB.WithContext(ctx).First(&studentORM, map[string]interface{}{
		"id": id,
	})
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &entities.Student{
		ID:   int64(studentORM.ID),
		Name: studentORM.Name,
		CPF:  studentORM.CPF,
		RG:   studentORM.RG,
	}, nil
}

func (s StudentRepository) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tx := db.DB.WithContext(ctx).Delete(&orm.Student{}, id)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s StudentRepository) Update(student *entities.Student) (*entities.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := map[string]interface{}{
		"id": student.ID,
	}
	tx := db.DB.WithContext(ctx).Where(query).UpdateColumns(&orm.Student{
		Name: student.Name,
		CPF:  student.CPF,
		RG:   student.RG,
	})
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if err := tx.Error; err != nil {
		return nil, err
	}
	return student, nil
}
