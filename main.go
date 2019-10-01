package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File(".ui/dist/ui/index.html")
		} else {
			c.File(".ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.GET("/file", handlers.GetFileListHandler)
	r.POST("/file", handlers.AddFileHandler)
	r.DELETE("/file/:id", handlers.DeleteFileHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
