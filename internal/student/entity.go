package student

import "github.com/labstack/echo/v4"

type Core struct {
	ID   int
	NPM  int // Nomor Induk Mahasiswa
	Name string
}

type Handler interface {
	Create(ctx echo.Context) error
	GetList(ctx echo.Context) error
	GetStudentByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type Service interface {
	Create(studentCore Core) error
	GetList() ([]Core, error)
	GetStudentByID(studentID int) (Core, error)
	Update(studentID int, studentCore Core) error
	Delete(studentID int) error
}

type Repository interface {
	Create(studentCore Core) error
	GetList() ([]Core, error)
	GetStudentByID(studentID int) (Core, error)
	Update(studentID int, studentCore Core) error
	Delete(studentID int) error
}
