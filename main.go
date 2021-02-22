package web


import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SayHi() {
	fmt.Println("Hello Golang!!!!")
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	SayHi()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"
}
