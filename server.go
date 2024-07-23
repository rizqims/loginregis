package main

import (
	"database/sql"
	"fmt"
	"log"
	"loginregis/config"
	"loginregis/controller"
	"loginregis/repository"
	"loginregis/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Server struct {
	uS      service.UserService
	g       *gin.Engine
	portApp string
}

func (s *Server) initiateServer() {
	group := s.g.Group("/api/v1")
	controller.NewUserController(group, s.uS).Route()
}

func (s *Server) Start() {
	s.initiateServer()
	s.g.Run(s.portApp)
}

func NewServer() *Server {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error on ENV: ")
	}

	db := &sql.DB{}
	if cfg.Driver == "mysql" {
		// your_username:your_password@tcp(127.0.0.1:3306)/your_database_name
		dbConn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
		db, err = sql.Open(cfg.Driver, dbConn)

	} else if cfg.Driver == "postgres" {
		dbConn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
		db, err = sql.Open(cfg.Driver, dbConn)
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	return &Server{
		uS:      userService,
		g:       gin.Default(),
		portApp: cfg.PortApp,
	}
}
