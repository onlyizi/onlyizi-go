package config

type Redis struct {
	Host     string
	Port     int
	Password string
}

func RedisConfig() Redis {

	return Redis{
		Host:     Get("REDIS_HOST", "localhost"),
		Port:     GetInt("REDIS_PORT", 6379),
		Password: Get("REDIS_PASSWORD", ""),
	}
}
