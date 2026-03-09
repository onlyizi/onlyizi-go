package config

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func PostgresConfig() Postgres {

	return Postgres{
		Host:     Get("POSTGRES_HOST", "localhost"),
		Port:     GetInt("POSTGRES_PORT", 5432),
		User:     Get("POSTGRES_USER", "postgres"),
		Password: Get("POSTGRES_PASSWORD", "postgres"),
		DB:       Get("POSTGRES_DB", "postgres"),
	}
}
