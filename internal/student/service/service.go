package service

import (
	"errors"
	"log"
	"wide_technologies/internal/student"
)

type studentService struct {
	repo student.Repository
}

func New(r student.Repository) student.Service {
	return &studentService{repo: r}
}

func (ss *studentService) GetStudentByID(studentID int) (student.Core, error) {
	res, err := ss.repo.GetStudentByID(studentID)
	if err != nil {
		log.Println(err)
		return student.Core{}, errors.New("internal server error")
	}
	return res, nil
}

func (ss *studentService) Create(studentCore student.Core) error {
	return nil
}
