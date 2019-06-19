package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/youtangai/cloud-training/model"
	"github.com/youtangai/cloud-training/service"
	"net/http"
)

type ISignController interface {
	SignIn(*gin.Context)
	SignUp(*gin.Context)
}

type SignController struct {
	srv service.ISignService
}

func NewSignController(srv service.ISignService) ISignController {
	return SignController{
		srv: srv,
	}
}

func (ctrl SignController) SignIn(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": fmt.Sprintf("failed to unmarshal json data. err=%s", err),
		})
		return
	}

	token, err := ctrl.srv.GetAccessToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("failed to find accesstoken. err=%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
	return
}

func (ctrl SignController) SignUp(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": fmt.Sprintf("failed to unmarshal json data. err=%s", err),
		})
		return
	}

	err = ctrl.srv.SignUpUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.Status(http.StatusCreated)
	return
}