package repository

import (
	"fmt"
	"wide_technologies/internal/student"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) student.Service {
	return &studentRepository{db: db}
}

func (ss *studentRepository) GetStudentByID(studentID int) (student.Core, error) {
	model := Student{}
	tx := ss.db.First(&model, studentID)
	if tx.Error != nil {
		return student.Core{}, fmt.Errorf("get student in student repository error: %w", tx.Error)
	}
	core := student.Core{
		ID:   int(model.ID),
		NPM:  model.NPM,
		Name: model.Name,
	}
	return core, nil
}

func (ss *studentRepository) Create(studentCore student.Core) error {
	model := Student{
		NPM:  studentCore.NPM,
		Name: studentCore.Name,
	}

	tx := ss.db.Create(&model)
	if tx.Error != nil {
		return fmt.Errorf("create student in student repository error: %w", tx.Error)
	}

	return nil
}
