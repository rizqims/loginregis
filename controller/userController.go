package controller

import (
	"fmt"
	"loginregis/model"
	"loginregis/model/dto"
	"loginregis/service"
	"loginregis/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	rg          *gin.RouterGroup
	userService service.UserService
}

func (u *UserController) Route() {
	group := u.rg.Group("/users")
	group.POST("/regiter", u.RegisHandler)
	group.POST("/login", u.LoginHandler)
}

func (u *UserController) RegisHandler(c *gin.Context) {
	var payload model.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrorRes(c, http.StatusBadRequest, "request is invalid!")
	}

	response, err := u.userService.Register(payload)
	if err != nil {
		errory := fmt.Errorf("server error: %v", err)
		util.SendErrorRes(c, http.StatusInternalServerError, errory.Error())
	}

	util.SendSuccessRes(c, http.StatusOK, "Register Success", response)
}

func (u *UserController) LoginHandler(c *gin.Context) {
	var payload dto.LoginDto
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrorRes(c, http.StatusBadRequest, "request is invalid!")
	}

	err = u.userService.Login(payload)
	if err != nil {
		errory := fmt.Errorf("server error: %v", err)
		util.SendErrorRes(c, http.StatusInternalServerError, errory.Error())
	}

	util.SendSuccessRes(c, http.StatusOK, "Login Success", 0)

}

func NewUserController(rg *gin.RouterGroup, uS service.UserService) *UserController {
	return &UserController{
		rg:          rg,
		userService: uS,
	}
}
