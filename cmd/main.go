package main

import (
	"github.com/puper/go-app"
	"github.com/puper/go-app/handlers"
)

func main() {
	goapp.AppInstance.Engine().GET("/test/:id", handlers.Index)
	goapp.AppInstance.Run()
}
