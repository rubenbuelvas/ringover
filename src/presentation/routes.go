package presentation

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, taskHandler TaskHandler) {
	tasksGroup := router.Group("/tasks")
	tasksGroup.GET("", taskHandler.GetTasksHandler)
	tasksGroup.GET("/:id/subtasks", taskHandler.GetSubtasksHandler)
	tasksGroup.POST("", taskHandler.CreateTaskHandler)
	tasksGroup.PATCH("/:id", taskHandler.PatchTaskHandler)
	tasksGroup.DELETE("/:id", taskHandler.DeleteTaskHandler)
}
