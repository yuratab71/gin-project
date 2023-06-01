package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Body struct {
	Name string
}

func main() {
	godotenv.Load(".env")
	engine := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))
	engine.GET("/", func(c *gin.Context) {
		body := Body{"test"}
		c.JSON(http.StatusAccepted, &body)
	})
	engine.Run(os.Getenv("PORT"))
}
