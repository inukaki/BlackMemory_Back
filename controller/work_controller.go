package controller

import (
	"go_rest_api/model"
	"go_rest_api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IWorkController interface {
	CreateWork(c echo.Context) error
	UpdateWork(c echo.Context) error
	GetWorkByDate(c echo.Context) error
	GetAllWorks(c echo.Context) error
}

type workController struct {
	wu usecase.IWorkUsecase
}

func NewWorkController(wu usecase.IWorkUsecase) IWorkController {
	return &workController{wu}
}

func (wc *workController) CreateWork(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	work := model.Work{}
	if err := c.Bind(&work); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	work.UserID = uint(userId.(float64))
	taskRes, err := wc.wu.CreateWork(work)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (wc *workController) UpdateWork(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	workDate := c.Param("workDate")
	// workId, _ := strconv.Atoi(id)

	work := model.Work{}
	if err := c.Bind(&work); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// work.UserID = uint(userId.(float64))
	workRes, err := wc.wu.UpdateWork(work, uint(userId.(float64)), workDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workRes)
}

func (wc *workController) GetWorkByDate(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	workDate := c.Param("workDate")
	workRes, err := wc.wu.GetWorkByDate(uint(userId.(float64)), string(workDate))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workRes)
}

func (wc *workController) GetAllWorks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	worksRes, err := wc.wu.GetAllWorks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, worksRes)
}
