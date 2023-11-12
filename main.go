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
	// taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	workRepository := repository.NewWorkRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	workUseCase := usecase.NewWorkUseCase(workRepository)
	// taskUseCase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUseCase)
	workController := controller.NewWorkController(workUseCase)
	// taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, workController)
	e.Logger.Fatal(e.Start(":8080"))
}
