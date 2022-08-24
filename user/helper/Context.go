package helper

import (
	"context"
	"time"
)

func NewDbContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
