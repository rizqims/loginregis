package service

import (
	"fmt"
	"loginregis/model"
	"loginregis/model/dto"
	"loginregis/repository"
	"loginregis/util"
)

type UserService interface {
	Register(payload model.User) (model.User, error)
	Login(payload dto.LoginDto) error
}

type userService struct {
	userRepo repository.UserRepo
}

func (u *userService) Register(payload model.User) (model.User, error) {
	hashedPass, err := util.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("hashError: %v", err)
	}

	payload.Password = hashedPass

	newUser, err := u.userRepo.Register(payload)
	if err != nil {
		return model.User{}, fmt.Errorf("GetRepo error: %v", err)
	}

	return newUser, nil
}

func (u *userService) Login(payload dto.LoginDto) error {
	user, err := u.userRepo.Login(payload)
	if err != nil {
		return fmt.Errorf("GetRepo error: %v", err)
	}

	err = util.CheckHash(payload.Password, user.Password)
	if err != nil {
		return fmt.Errorf("Login error: checkhash error: %v", err)
	}
	return nil
}

func NewUserService(uR repository.UserRepo) UserService {
	return &userService{
		userRepo: uR,
	}
}
