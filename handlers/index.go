package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/puper/go-app"
)

func Index(c *gin.Context) {
	var one struct {
		Id  int
		Num string
	}
	goapp.AppInstance.GetDB("default").Read().Select("id", "num").From("test").Where(&dbx.HashExp{"id": c.Param("id")}).One(&one)
	c.JSON(200, one)
}
