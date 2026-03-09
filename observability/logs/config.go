package logs

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type Config struct {
	Service     string
	Environment Environment
	Version     string
	Level       string
}
