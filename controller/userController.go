package controller

import (
	"fmt"
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
	group.POST("/register", u.RegisHandler)
	group.POST("/login", u.LoginHandler)
}

func (u *UserController) RegisHandler(c *gin.Context) {
	var payload dto.RegisDto
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrorRes(c, http.StatusBadRequest, "request is invalid!")
		return
	}

	response, err := u.userService.Register(payload)
	if err != nil {
		errory := fmt.Errorf("server error: %v", err)
		util.SendErrorRes(c, http.StatusInternalServerError, errory.Error())
		return
	}

	util.SendSuccessRes(c, http.StatusOK, "Register Success", response)
}

func (u *UserController) LoginHandler(c *gin.Context) {
	var payload dto.LoginDto
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrorRes(c, http.StatusBadRequest, "request is invalid!")
		return
	}

	err = u.userService.Login(payload)
	if err != nil {
		errory := fmt.Errorf("server error: %v", err)
		util.SendErrorRes(c, http.StatusInternalServerError, errory.Error())
		return
	}

	util.SendSuccessRes(c, http.StatusOK, "Login Success", nil)

}

func NewUserController(rg *gin.RouterGroup, uS service.UserService) *UserController {
	return &UserController{
		rg:          rg,
		userService: uS,
	}
}
