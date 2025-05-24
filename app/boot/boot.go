package boot

import (
	"github.com/goyalaadi12071998/app/controllers"
	"github.com/goyalaadi12071998/app/services"
)

func Init() {

	// Added a dependency injection from controller => service and we can extent to db or third party requests as well
	userService := services.InitialiseUserService()
	controllers.InitialiseUserController(userService)

	taskService := services.InitialiseTaskService()
	controllers.InitaliseTaskController(taskService)
}
