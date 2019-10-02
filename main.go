package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/xrustalik/protected-notebook/handlers"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTION, POST, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,origin,Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTION" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
