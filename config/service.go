package config

type Service struct {
	Name        string
	Version     string
	Environment string
	Hostname    string
}

func ServiceConfig() Service {

	return Service{
		Name:        Get("SERVICE_NAME", "onlyizi-service"),
		Version:     Get("SERVICE_VERSION", "0.1.0"),
		Environment: Get("ENVIRONMENT", "development"),
		Hostname:    Get("HOSTNAME", ""),
	}
}
