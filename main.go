package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/youtangai/cloud-training/config"
	"github.com/youtangai/cloud-training/controller"
	"github.com/youtangai/cloud-training/model"
	"github.com/youtangai/cloud-training/repository"
	"github.com/youtangai/cloud-training/service"
	"log"
	"net/http"
)

func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatalf("failed initialize db. err=%s", err)
	}

	defer db.Close()

	ctrl, err := initializeController(db)
	if err != nil {
		log.Fatalf("failed initialize sign controller. err=%s", err)
	}

	router := setupRouter(*ctrl)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("failed run gin server:", err)
	}
}

func initializeController(db *gorm.DB) (*controller.ISignController, error) {
	repo := repository.NewUserRepo(db)
	srv := service.NewSignService(repo)
	ctrl := controller.NewSignController(srv)
	return &ctrl, nil
}

func initializeDB() (*gorm.DB, error) {
	connectionString := config.GetConnectionString()
	db ,err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to establish mysql connection. connectionString=%s, err=%s", connectionString, err)
	}

	db.AutoMigrate(&model.User{})

	return db, nil
}

func setupRouter(ctrl controller.ISignController) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.POST("/signin", ctrl.SignIn)
	router.POST("/signup", ctrl.SignUp)

	return router
}
