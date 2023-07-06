package logging

import (
	"context"
)

type DummyLogger struct {
}

func (l *DummyLogger) HandledError(ctx context.Context, err error, metadata Metadata) {
	// dummy
}

func (l *DummyLogger) StartSession(ctx context.Context) context.Context {
	return context.TODO()
}

func (l *DummyLogger) Info(s string, metadata Metadata) {
	// dummy
}

func (l *DummyLogger) Error(ctx context.Context, err error, metadata Metadata) {
	// dummy
}

func (l *DummyLogger) Fatal(ctx context.Context, err error, metadata Metadata) {
	// dummy
}

func (l *DummyLogger) Sync() {
	// dummy
}
