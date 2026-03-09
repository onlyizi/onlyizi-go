package config

type HTTP struct {
	Port int
}

func HTTPConfig() HTTP {

	return HTTP{
		Port: GetInt("HTTP_PORT", 8080),
	}
}
