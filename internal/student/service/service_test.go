package service

import (
	"testing"
	"wide_technologies/internal/student"
	"wide_technologies/internal/student/repository"

	"github.com/stretchr/testify/assert"
)

func TestGetStudentByID(t *testing.T) {
	sampleData := student.Core{
		ID:   1,
		NPM:  54418672,
		Name: "Helmi",
	}

	t.Run("should return list catalogues data", func(t *testing.T) {
		repoMock := new(repository.StudentRepositoryMock)
		repoMock.On("GetStudentByID", sampleData.ID).Return(sampleData, nil).Once()

		srv := New(repoMock)
		result, err := srv.GetStudentByID(sampleData.ID)

		assert.Nil(t, err)
		assert.Equal(t, sampleData.Name, result.Name)
	})
}
