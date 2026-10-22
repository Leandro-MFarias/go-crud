package controller

import (
	"fmt"

	"github.com/Leandro-MFarias/go-crud/src/config/validation"
	"github.com/Leandro-MFarias/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println("Erro na validação:", err.Error())

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
