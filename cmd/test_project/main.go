package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcusbianchi/test_project/internal/routers"
)

func main() {
	r := gin.New()

	r = routers.SetupRouter(r)
	r.Use(gin.Recovery())
	if err := r.Run(); err != nil {
		panic(err)
	}
}
