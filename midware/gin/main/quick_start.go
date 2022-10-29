package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Group()
	r.GET("/person/eat")
	r.GET("/person/laugh")
	r.GET("/person/lying")

	r.Handle("GET", "/ping", func(c *gin.Context) {
		var name string
		err := c.ShouldBind(&name)
		if err != nil {
			fmt.Println("should bind query error")
			return
		}
		fmt.Printf("name is %s", name)
		return
	})
	r.Run("127.0.0.1:8001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
