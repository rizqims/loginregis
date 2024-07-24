package service

import (
	"fmt"
	"loginregis/model"
	"loginregis/model/dto"
	"loginregis/repository"
	"loginregis/util"
	"strconv"
)

type UserService interface {
	Register(payload dto.RegisDto) (model.User, error)
	Login(payload dto.LoginDto) error
}

type userService struct {
	userRepo repository.UserRepo
}

func (u *userService) Register(payload dto.RegisDto) (model.User, error) {
	if payload.Password != payload.PassConfirm {
		return model.User{}, fmt.Errorf("PassConfirm error: password and confirmPass invalid")
	}

	hashedPass, err := util.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("hashError: %v", err)
	}

	payload.Password = hashedPass
	age, err := strconv.Atoi(payload.Age)
	if err != nil {
		return model.User{}, fmt.Errorf("strconv error: %v", err)
	}

	user := model.User{
		Name:     payload.Name,
		Address:  payload.Address,
		Age:      age,
		Username: payload.Username,
		Password: payload.Password,
	}

	newUser, err := u.userRepo.Register(user)
	if err != nil {
		return model.User{}, fmt.Errorf("GetRepo error: %v", err)
	}

	return newUser, nil
}

func (u *userService) Login(payload dto.LoginDto) error {
	if payload.Username == "" || payload.Password == "" {
		return fmt.Errorf("service error: either username or password must be inserted!")
	}

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
