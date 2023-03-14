package handler

import (
	"strconv"
	"wide_technologies/internal/student"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	srv student.Service
}

func New(s student.Service) student.Handler {
	return &studentHandler{srv: s}
}

func (sh *studentHandler) GetStudentByID(ctx echo.Context) error {
	str := ctx.Param("student_id")
	studentID, _ := strconv.Atoi(str)
	res, err := sh.srv.GetStudentByID(studentID)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": err})
	}
	response := studentReponse{
		ID:   res.ID,
		NPM:  res.NPM,
		Name: res.Name,
	}
	webReponse := map[string]interface{}{
		"message": "success",
		"data":    response,
	}
	return ctx.JSON(200, webReponse)
}

func (sh *studentHandler) Create(ctx echo.Context) error {
	request := studentRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": err})
	}
	studentCore := student.Core{
		NPM:  request.NPM,
		Name: request.Name,
	}
	if err := sh.srv.Create(studentCore); err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": err})
	}
	return ctx.JSON(201, map[string]interface{}{"message": "success"})
}
