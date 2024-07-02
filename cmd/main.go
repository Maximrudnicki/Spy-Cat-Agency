package main

import (
	"fmt"
	"log"
	"net/http"
	"test_rudnytskyi/cmd/config"
	"test_rudnytskyi/cmd/controllers"
	"test_rudnytskyi/cmd/models"
	repo "test_rudnytskyi/cmd/repositories"
	"test_rudnytskyi/cmd/router"
	"test_rudnytskyi/cmd/services"

	"github.com/rs/cors"
)

func main() {
	// env variables
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Database (PostgreSQL)
	db := config.ConnectionDB(&loadConfig)
	db.Table("cats").AutoMigrate(&models.Cat{})
	db.Table("missions").AutoMigrate(&models.Mission{})
	db.Table("targets").AutoMigrate(&models.Target{})

	// repos
	catRepository := repo.NewCatRepositoryImpl(db)
	missionRepository := repo.NewMissionRepositoryImpl(db)

	// services
	catService := services.NewCatServiceImpl(catRepository)
	missionService := services.NewMissionServiceImpl(missionRepository)

	// controllers
	catController := controllers.NewCatController(catService)
	missionController := controllers.NewMissionController(missionService)

	r := router.NewRouter(catController, missionController)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", loadConfig.PORT),
		Handler: handler,
	}

	log.Printf("Listening at %s\n", loadConfig.PORT)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
