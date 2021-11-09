package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blogpoc/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("blogPOC")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	port := viper.GetString("app.port")

	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	//Add cors headers
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	// Route => handler
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blog API!\n")
	})
	blogGroup := app.Group("/v1/api")
	blogGroup.Use(middleware.JWT([]byte(viper.GetString("jwt.key"))))
	app.POST("/login", controllers.Login)
	blogGroup.GET("/blog/:id", controllers.GetBlogs)
	blogGroup.POST("/blog", controllers.PostBlog)
	// Start server
	app.Logger.Fatal((app.Start((":" + port))))
}
