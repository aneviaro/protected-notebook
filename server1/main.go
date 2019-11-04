package main

import (
	"path"
	"path/filepath"

	"protected-notebook/server1/handlers"

	"github.com/gin-gonic/gin"
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
	r.POST("/login", handlers.Login)
	r.POST("/rsa", handlers.SetRSAPublicKey)
	r.GET("/file/:name/:client", handlers.GetFileContent)
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
