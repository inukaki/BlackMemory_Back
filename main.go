package main

import (
	"go_rest_api/controller"
	"go_rest_api/db"
	"go_rest_api/repository"
	"go_rest_api/router"
	"go_rest_api/usecase"
	"go_rest_api/validator"
	"os"
)

func main() {
	os.Setenv("GO_ENV", "dev")

	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	taskUseCase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
