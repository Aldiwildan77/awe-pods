package main

import (
	"awepods/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Viper Setup
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s ", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Routing Mux Setup
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("ini root")
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), router); err != nil {
		panic("Can't run your server")
	}

}
