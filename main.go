package main

import (
	"log"
	"wide_technologies/config"
	_studentHandler "wide_technologies/internal/student/handler"
	_studentRepository "wide_technologies/internal/student/repository"
	_studentService "wide_technologies/internal/student/service"
	"wide_technologies/pkg/database"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := database.InitMysql(conf)

	studentRepo := _studentRepository.New(db)
	studentSrv := _studentService.New(studentRepo)
	studentHandler := _studentHandler.New(studentSrv)

	e := echo.New()
	e.POST("/students", studentHandler.Create)
	e.GET("/students/:student_id", studentHandler.GetStudentByID)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
