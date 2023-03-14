package repository

import (
	"wide_technologies/internal/student"

	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (srm *StudentRepositoryMock) GetStudentByID(studentID int) (student.Core, error) {
	args := srm.Mock.Called(studentID)
	ret0, _ := args[0].(student.Core)
	ret1, _ := args[1].(error)
	return ret0, ret1
}

func (srm *StudentRepositoryMock) Create(studentCore student.Core) error {
	args := srm.Mock.Called(studentCore)
	ret1, _ := args[1].(error)
	return ret1
}

// func (srm *StudentRepositoryMock) GetStudentByID(studentID int) (student.Core, error) {
// args := srm.Mock.Called(studentID)

// var ret0 student.Core
// 	rf0, ok := args.Get(0).(func(int) student.Core)
// 	if ok {
// 		ret0 = rf0(studentID)
// 	} else {
// 		ret0 = args.Get(0).(student.Core)
// 	}

// 	var ret1 error
// 	rf1, ok := args.Get(1).(func(int) error)
// 	if ok {
// 		ret1 = rf1(studentID)
// 	} else {
// 		ret1 = args.Error(1)
// 	}

// }

func (srm *StudentRepositoryMock) GetList() ([]student.Core, error) {

	return nil, nil
}

func (srm *StudentRepositoryMock) Update(studentID int, studentCore student.Core) error {

	return nil
}

func (srm *StudentRepositoryMock) Delete(studentID int) error {

	return nil
}
