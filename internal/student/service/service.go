package service

import (
	"fmt"
	"log"
	"regexp"
	"wide_technologies/internal/student"
	"wide_technologies/pkg/util"
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
		return student.Core{}, util.SetError(500, "internal server error")
	}
	return res, nil
}

func (ss *studentService) Create(studentCore student.Core) error {
	if len(studentCore.Name) < 3 {
		log.Println("sutdent name must greater than 3")
		return util.SetError(400, "bad request")
	}
	// Check the npm input using regex
	// "^" matches the start of the string.
	// "[0-9]" Match any digits
	// "{4,}" matches the previous character (which is [0-9]) 4 or more times.
	// "$" matches the previous character (which is [0-9]) 4 or more times.
	regex := regexp.MustCompile(`^[0-9]{4,}$`)
	isMatch := regex.MatchString(fmt.Sprintf("%d", studentCore.NPM))
	if !isMatch {
		log.Println("student npm must be integer and greater than 4 digit")
		return util.SetError(400, "bad request")
	}
	if err := ss.repo.Create(studentCore); err != nil {
		log.Println(err)
		return util.SetError(500, "internal server error")
	}
	return nil
}

func (ss *studentService) GetList() ([]student.Core, error) {
	res, err := ss.repo.GetList()
	if err != nil {
		log.Println(err)
		return nil, util.SetError(500, "internal server error")
	}
	return res, nil
}

func (ss *studentService) Update(studentID int, studentCore student.Core) error {
	if err := ss.repo.Update(studentID, studentCore); err != nil {
		log.Println(err)
		return util.SetError(500, "internal server error")
	}
	return nil
}

func (ss *studentService) Delete(studentID int) error {
	if err := ss.repo.Delete(studentID); err != nil {
		log.Println(err)
		return util.SetError(500, "internal server error")
	}
	return nil
}
