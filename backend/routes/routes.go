package routes

import (
	"github.com/1234arushi/JD-Analyzer/controller"
	"github.com/gin-gonic/gin"
)

func LoadServices(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	{
		v1.POST("/analyze", controller.SendResponse)
	}
}
