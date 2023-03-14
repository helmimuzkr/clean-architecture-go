package handler

import (
	"log"
	"strconv"
	"wide_technologies/internal/student"
	"wide_technologies/pkg/util"

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
		log.Println(err)
		return ctx.JSON(util.ErrorResponse(err))
	}
	studentCore := student.Core{
		NPM:  request.NPM,
		Name: request.Name,
	}
	if err := sh.srv.Create(studentCore); err != nil {
		return ctx.JSON(util.ErrorResponse(err))
	}
	return ctx.JSON(util.SuccessResponse(201, "success"))
}

func (sh *studentHandler) GetList(ctx echo.Context) error {
	res, err := sh.srv.GetList()
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{"message": err})
	}
	dataResponses := []studentReponse{}
	for _, v := range res {
		r := studentReponse{
			ID:   v.ID,
			NPM:  v.NPM,
			Name: v.Name,
		}
		dataResponses = append(dataResponses, r)
	}
	return ctx.JSON(util.SuccessResponse(200, "success", dataResponses))
}

func (sh *studentHandler) Update(ctx echo.Context) error {
	strID := ctx.Param("student_id")
	studentID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		return ctx.JSON(util.ErrorResponse(err))
	}
	request := studentRequest{}
	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(util.ErrorResponse(err))
	}
	studentCore := student.Core{
		NPM:  request.NPM,
		Name: request.Name,
	}
	if err := sh.srv.Update(studentID, studentCore); err != nil {
		return ctx.JSON(util.ErrorResponse(err))
	}
	return ctx.JSON(util.SuccessResponse(200, "success"))
}

func (sh *studentHandler) Delete(ctx echo.Context) error {
	strID := ctx.Param("student_id")
	studentID, err := strconv.Atoi(strID)
	if err != nil {
		return ctx.JSON(util.ErrorResponse(err))
	}
	if err := sh.srv.Delete(studentID); err != nil {
		return ctx.JSON(util.ErrorResponse(err))
	}
	return ctx.JSON(util.SuccessResponse(200, "success"))
}
