package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainRoutes() *gin.Engine {
	route := gin.Default()
	route.LoadHTMLGlob("templates/**/*.html")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	route.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	route.GET("/customer/:id/", func(c *gin.Context) {
		id := c.Param("id")
		result, ok := customer[id]
		fmt.Println("test", result)

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "index.html",
			map[string]interface{}{
				"data": result,
			})
	})

	route.Static("/public", "./public")
	// r.StaticFS("/public", http.Dir("./public"))
	return route
}
