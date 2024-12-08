package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"

	"github.com/bianavic/fullcycle_apis/configs"
	"github.com/bianavic/fullcycle_apis/internal/entity"
	"github.com/bianavic/fullcycle_apis/internal/infra/database"
	"github.com/bianavic/fullcycle_apis/internal/infra/webserver/handlers"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	//http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8080", r)
}
