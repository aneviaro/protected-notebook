package handlers

import (
	"net/http"

	"github.com/xrustalik/text-editor-backend/file"

	"github.com/gin-gonic/gin"
)

func GetFileListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, file.Get())
}
