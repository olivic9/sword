package logging

import (
	"context"
	"sword-project/pkg/zap"

	uberZap "go.uber.org/zap"
)

type ApplicationLogger struct {
	logger *uberZap.SugaredLogger
}

var Logger GenericLogger = &DummyLogger{}

func InitializeApplicationLogger() {

	Logger = &ApplicationLogger{
		logger: zap.NewZapSugaredLogger(),
	}
}
func (l *ApplicationLogger) StartSession(ctx context.Context) context.Context {
	return ctx
}

func (l *ApplicationLogger) Info(s string, metadata Metadata) {
	l.logger.Infow(s, metadata.ToZapMetadata()...)
}

func (l *ApplicationLogger) Error(ctx context.Context, err error, metadata Metadata) {
	l.logger.Errorw(err.Error(), metadata.ToZapMetadata()...)
}

func (l *ApplicationLogger) Fatal(ctx context.Context, err error, metadata Metadata) {
	l.logger.Fatalw(err.Error(), metadata.ToZapMetadata()...)
}

func (l *ApplicationLogger) Sync() {
	_ = l.logger.Sync()
}
