package common

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	opentracingConfig := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
			SamplingServerURL:        "",
			SamplingRefreshInterval:  0,
			MaxOperations:            0,
			OperationNameLateBinding: false,
			Options:                  nil,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	return opentracingConfig.NewTracer()
}
