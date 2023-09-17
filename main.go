package main

import (
	"todo-app/controller"
	"todo-app/db"
	"todo-app/repository"
	"todo-app/router"
	"todo-app/usecase"
	"todo-app/validator"
)

func main() {
	db := db.NewDb()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
