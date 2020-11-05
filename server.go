package main

import (
	"awepods/config"
	"awepods/database"
	users "awepods/user"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var (
	configuration config.Configurations
	postgres      database.PostgresDB
)

func init() {
	loadConfigurations()

	postgres = database.PostgresDB{}
}

func main() {
	// Database Setup
	pgInstance := postgres.NewInstance(&configuration.Database.PostgresDB)
	pgInstance.LogMode(true)

	defer pgInstance.Close()

	// Routing Mux Setup
	router := mux.NewRouter()

	// Register Modules
	userRepository := users.NewUserRepository(pgInstance)
	userService := users.NewUserService(&userRepository)
	userController := users.NewUserController(&userService)
	userController.Route(router)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()

	// Server Start
	addr := fmt.Sprintf("%s:%d", configuration.Server.Host, configuration.Server.Port)

	fmt.Printf("Server is running at http://%s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}

func loadConfigurations() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s ", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
