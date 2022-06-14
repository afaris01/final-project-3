package router

import (
	"final-project-3/controllers"
	"final-project-3/middlewares"

	"final-project-3/database"
	"final-project-3/repositories"
	"final-project-3/services"

	"github.com/gin-gonic/gin"
)

func MulaiApp() *gin.Engine {
	db := database.MulaiDB()

	// repository
	userRepository := repositories.NewUserRepository(db)
	taskRepository := repositories.NewTaskRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)

	// service
	userService := services.NewUserService(userRepository)
	taskService := services.NewTaskService(taskRepository)
	categoryService := services.NewCategoryService(categoryRepository)

	// controller
	userController := controllers.NewUserController(userService)
	taskController := controllers.NewTaskController(taskService)
	categoryController := controllers.NewCategoryController(categoryService, userService)

	router := gin.Default()

	// routing

	// user
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.Login)
		userRouter.PUT("/update-account", middlewares.AuthMiddleware(), userController.UpdateUser)
		userRouter.DELETE("/delete-account", middlewares.AuthMiddleware(), userController.DeleteUser)
	}

	// task
	taskRouter := router.Group("/tasks")
	{
		taskRouter.POST("/", middlewares.AuthMiddleware(), taskController.CreateNewTask)
		taskRouter.GET("/", middlewares.AuthMiddleware(), taskController.GetAllTask)
		taskRouter.PUT("/:id", middlewares.AuthMiddleware(), taskController.UpdateTask)
		taskRouter.PATCH("/update-status/:id", middlewares.AuthMiddleware(), taskController.UpdateStatusTask)
		taskRouter.PATCH("/update-category/:id", middlewares.AuthMiddleware(), taskController.UpdateCategoryTask)
		taskRouter.DELETE("/:id", middlewares.AuthMiddleware(), taskController.DeleteTask)
	}

	// category
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.POST("/", middlewares.AuthMiddleware(), categoryController.CreateCategory)
		categoryRouter.GET("/", middlewares.AuthMiddleware(), categoryController.GetAllCategory)
		categoryRouter.PATCH("/:id", middlewares.AuthMiddleware(), categoryController.UpdateCategory)
		categoryRouter.DELETE("/:id", middlewares.AuthMiddleware(), categoryController.DeleteCategory)
	}
	return router
}
