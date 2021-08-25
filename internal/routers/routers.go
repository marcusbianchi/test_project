package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcusbianchi/test_project/internal/controllers"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	//API route for version 1
	v1 := r.Group("/api/v1")

	v1.POST("/album", controllers.AddAlbum)
	return r
}
