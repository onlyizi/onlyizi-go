package app

import "context"

type Service interface {
	Name() string
	Start() error
	Shutdown(ctx context.Context) error
}
