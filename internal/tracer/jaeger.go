package tracer

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/yvv4git/eger/internal/config"
	"io"
	"log"
)

// InitTracer - used for init jaeger
func InitTracer(cfg config.JaegerConf) io.Closer {
	cfgJ := jaegercfg.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeProbabilistic,
			Param: cfg.SamplingPercentage,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: fmt.Sprintf("%s:%s", cfg.AgentHost, cfg.AgentPort),
		},
	}

	tracer, closer, err := cfgJ.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)

	if err != nil {
		log.Fatalln("couldn't create jaeger tracer")
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}
