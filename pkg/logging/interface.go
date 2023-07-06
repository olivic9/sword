package logging

import (
	"context"
)

type GenericLogger interface {
	StartSession(ctx context.Context) context.Context
	Info(s string, metadata Metadata)
	Error(ctx context.Context, err error, metadata Metadata)
	Fatal(ctx context.Context, err error, metadata Metadata)
	Sync()
}
