package config

type Configurations struct {
  Server   ServerConfigurations
  Database DatabaseConfigurations
}

type ServerConfigurations struct {
  Host string
  Port int
}

type DatabaseConfigurations struct {
  PostgresDB DatabaseConfiguration
}

type DatabaseConfiguration struct {
  DBHost     string
  DBPort     int
  DBName     string
  DBUser     string
  DBPassword string
}
