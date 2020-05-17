package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func handleServer() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hai")
	})

	e.GET("/user", AllUsers)                //Get All User
	e.GET("/user/:name", ShowUser)          //Get User with Param
	e.POST("/user/:name/:email", NewUser)   // create user
	e.PUT("/user/:name/:email", UpdateUser) // update all data user
	e.DELETE("/user/:name", DeleteUser)     // delete user

	e.Logger.Fatal(e.Start(":8000")) //start web server
}

func main() {

	InitialMigration()

	handleServer()
}
