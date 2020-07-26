package config

type Configurations struct {
	Server ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	PostgresDB PostgresDatabaseConfig
	MongoDB MongoDatabaseConfig
}

type PostgresDatabaseConfig struct {
	DBHost string
	DBPort int
	DBName string
	DBUser string
	DBPassword string
}

type MongoDatabaseConfig struct {
	DBHost string
	DBPort int
	DBName string
	DBUser string
	DBPassword string
}