package database

import (
  "awepods/config"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "log"
)

type PostgresDB struct {
  db *gorm.DB
}

func (pg *PostgresDB) NewInstance(conf *config.DatabaseConfiguration) *gorm.DB {
  args := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
    conf.DBHost,
    conf.DBPort,
    conf.DBUser,
    conf.DBName,
    conf.DBPassword,
    "disable")

  instance, err := gorm.Open("postgres", args)
  if err != nil {
    log.Fatal(err)
  }

  return instance
}
