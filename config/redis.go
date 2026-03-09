package config

type Redis struct {
	Host string
	Port int
}

func RedisConfig() Redis {

	return Redis{
		Host: Get("REDIS_HOST", "localhost"),
		Port: GetInt("REDIS_PORT", 6379),
	}
}
