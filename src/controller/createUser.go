package controller

import (
	"github.com/Leandro-MFarias/go-crud/src/config/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota errada")

	c.JSON(err.Code, err)
}
