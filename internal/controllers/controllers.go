package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcusbianchi/test_project/internal/models"
	"net/http"
)

func AddAlbum(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}
