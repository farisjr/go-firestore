package main

import (
	"fmt"
	"go-firestore/configs"
	"go-firestore/routes"
	"go-firestore/user/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//configs.InitDb()
	configs.InitPort()
	controller.AddUser()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configs.HTTP_PORT)))
}
