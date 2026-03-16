package postgres

import "context"

func Health(ctx context.Context) error {
	return DB().PingContext(ctx)
}
