package repository

import (
	"database/sql"
	"fmt"
	"loginregis/config"
	"loginregis/model"
	"loginregis/model/dto"
)

type UserRepo interface {
	Register(payload model.User) (model.User, error)
	Login(payload dto.LoginDto) (dto.LoginDto, error)
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) Register(payload model.User) (model.User, error) {
	t, err := u.db.Begin()
	if err != nil {
		return model.User{}, fmt.Errorf("begin error: %v", err)
	}

	var str, selectID string
	c, err := config.NewConfig()
	if c.Driver == "postgres" {
		selectID = "SELECT id FROM users WHERE username=$1"
		str = "INSERT INTO users (name, address, age, username, password) VALUES ($1,$2,$3,$4,$5)"
	} else if c.Driver == "mysql" {
		selectID = "SELECT id FROM users WHERE username=?"
		str = "INSERT INTO users (name, address, age, username, password) VALUES (?,?,?,?,?)"
	}

	newUser := model.User{}
	_, err = t.Exec(str,
		payload.Name,
		payload.Address,
		payload.Age,
		payload.Username,
		payload.Password,
	)
	if err != nil {
		return model.User{}, t.Rollback()
	}

	err = t.QueryRow(selectID, payload.Username).Scan(&newUser.ID)
	if err != nil {
		return model.User{}, t.Rollback()
	}

	err = t.Commit()
	if err != nil {
		return model.User{}, fmt.Errorf("commit error: %v", err)
	}

	newUser.Name = payload.Name
	newUser.Address = payload.Address
	newUser.Age = payload.Age
	newUser.Username = payload.Username
	newUser.Password = payload.Password

	return newUser, nil
}

func (u *userRepo) Login(payload dto.LoginDto) (dto.LoginDto, error) {
	var user = dto.LoginDto{}

	var logQuery string
	c, err := config.NewConfig()
	if c.Driver == "postgres" {
		logQuery = "SELECT password FROM users WHERE username=$1"
	} else if c.Driver == "mysql" {
		logQuery = "SELECT password FROM users WHERE username=?"
	}

	err = u.db.QueryRow(logQuery, payload.Username).Scan(
		&user.Password,
	)
	if err != nil {
		return dto.LoginDto{}, fmt.Errorf("login repo error: %v", err)
	}

	return user, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
