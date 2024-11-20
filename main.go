package main

import (
	"asset-mapping/library/config"
	"asset-mapping/library/db"
	"asset-mapping/pkg/handlers/assets"
	"asset-mapping/pkg/handlers/dashboard"
	mappings "asset-mapping/pkg/handlers/mapping"
	"asset-mapping/pkg/handlers/user"
	"asset-mapping/pkg/middleware"
	"asset-mapping/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	defer db.Disconnect()

	database := db.Client.Database(config.DB_NAME)

	userRepo := repository.NewUserRepository(database, "users")
	assetRepo := repository.NewAssetsRepository(database, "assets")
	mappingRepo := repository.NewMappingRepository(database, "mappings")

	userHandler := user.Newhandler(userRepo)
	assetsHandler := assets.Newhandler(assetRepo)
	mappingHandler := mappings.Newhandler(mappingRepo, assetRepo)
	dashboardHandler := dashboard.Newhandler(userRepo, mappingRepo)

	r := mux.NewRouter()

	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	employeeGroup := r.PathPrefix("/employee").Subrouter()
	assetsGroup := r.PathPrefix("/asset").Subrouter()
	mappingGroup := r.PathPrefix("/mapping").Subrouter()
	dashboardGroup := r.PathPrefix("/dashboard").Subrouter()

	employeeGroup.Use(middleware.AuthMiddleware)
	assetsGroup.Use(middleware.AuthMiddleware)
	mappingGroup.Use(middleware.AuthMiddleware)
	dashboardGroup.Use(middleware.AuthMiddleware)

	employeeGroup.HandleFunc("", userHandler.Create).Methods("POST")
	employeeGroup.HandleFunc("", userHandler.Update).Methods("PUT")
	employeeGroup.HandleFunc("", userHandler.Get).Methods("GET")
	employeeGroup.HandleFunc("/{id}", userHandler.Delete).Methods("DELETE")
	employeeGroup.HandleFunc("/{id}", userHandler.GetById).Methods("GET")

	assetsGroup.HandleFunc("", assetsHandler.Create).Methods("POST")
	assetsGroup.HandleFunc("", assetsHandler.Update).Methods("PUT")
	assetsGroup.HandleFunc("", assetsHandler.Get).Methods("GET")
	assetsGroup.HandleFunc("/{id}", assetsHandler.Delete).Methods("DELETE")
	assetsGroup.HandleFunc("/{id}", assetsHandler.GetById).Methods("GET")

	mappingGroup.HandleFunc("", mappingHandler.Create).Methods("POST")
	mappingGroup.HandleFunc("/employee/{userId}", mappingHandler.Get).Methods("GET")
	mappingGroup.HandleFunc("/{id}", mappingHandler.Delete).Methods("DELETE")

	dashboardGroup.HandleFunc("", dashboardHandler.Get).Methods("GET")

	http.ListenAndServe(":"+config.PORT, r)
}
