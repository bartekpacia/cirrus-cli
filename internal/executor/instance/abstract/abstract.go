package abstract

import (
	"context"
	"github.com/cirruslabs/cirrus-cli/internal/executor/instance/runconfig"
	"github.com/cirruslabs/echelon"
	"go.opentelemetry.io/otel/attribute"
)

type Instance interface {
	Run(context.Context, *runconfig.RunConfig) error
	WorkingDirectory(projectDir string, dirtyMode bool) string
	Close(ctx context.Context) error
	Attributes() []attribute.KeyValue
}

type WarmableInstance interface {
	// Warmup can be optionally called in case of a persistent worker is configured to be warm
	Warmup(context.Context, string, map[string]string, *echelon.Logger) error
}
