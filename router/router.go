package router

import (
	"go_rest_api/controller"
	//"net/http"
	"os"

	// echojwt "github.com/golang-jwt/jwt/v4"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, wc controller.IWorkController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/signup", uc.SingUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	// t := e.Group("/tasks")
	// t.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:  []byte(os.Getenv("SECRET")),
	// 	TokenLookup: "cookie:token",
	// }))
	// t.GET("", tc.GetAllTasks)
	// t.GET("/:taskId", tc.GetTaskByID)
	// t.POST("", tc.CreateTask)
	// t.PUT("/:taskId", tc.UpdateTask)
	// t.DELETE("/:taskId", tc.DeleteTask)
	w := e.Group("/works")
	w.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	w.POST("", wc.CreateWork)
	w.PUT("/:workDate", wc.UpdateWork)
	w.GET("/:workDate", wc.GetWorkByDate)
	w.GET("", wc.GetAllWorks)

	return e
}
