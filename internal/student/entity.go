package student

import "github.com/labstack/echo/v4"

type Core struct {
	ID   int
	NPM  int // Nomor Induk Mahasiswa
	Name string
}

type Handler interface {
	GetStudentByID(ctx echo.Context) error
	Create(ctx echo.Context) error
}

type Service interface {
	GetStudentByID(studentID int) (Core, error)
	Create(studentCore Core) error
}

type Repository interface {
	GetStudentByID(studentID int) (Core, error)
	Create(studentCore Core) error
}
