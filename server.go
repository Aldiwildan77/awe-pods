package main

import (
  "awepods/config"
  "awepods/database"
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/spf13/viper"
)

var (
  configuration config.Configurations
  postgres      database.PostgresDB
)

func init() {
  // Viper Setup
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

  postgres = database.PostgresDB{}

}

func main() {
  // Database Setup
  pgInstance := postgres.NewInstance(&configuration.Database.PostgresDB)
  defer pgInstance.Close()

  // Routing Mux Setup
  router := mux.NewRouter()

  // Server Start
  addr := fmt.Sprintf("%s:%d",
    configuration.Server.Host,
    configuration.Server.Port)

  fmt.Printf("Server is running at http://%s\n", addr)
  log.Fatal(http.ListenAndServe(addr, router))

}
