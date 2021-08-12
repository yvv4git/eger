package config

const (
	webHost                  = "localhost"
	webPort                  = "8080"
	jaegerServiceName        = "example"
	jaegerAgentHost          = "localhost"
	jaegerAgentPort          = "6831"
	jaegerSamplingPercentage = 0.91
)

type (
	// Config - main config
	Config struct {
		WebSrv     WebSrv
		JaegerConf JaegerConf
	}

	// WebSrv - web server config
	WebSrv struct {
		Host string
		Port string
	}

	// JaegerConf - jaeger config
	JaegerConf struct {
		ServiceName        string
		AgentHost          string
		AgentPort          string
		SamplingPercentage float64
	}
)

// NewConfig - simple factory for create Config instance
func NewConfig() *Config {
	return &Config{
		WebSrv: WebSrv{
			Host: webHost,
			Port: webPort,
		},
		JaegerConf: JaegerConf{
			ServiceName:        jaegerServiceName,
			AgentHost:          jaegerAgentHost,
			AgentPort:          jaegerAgentPort,
			SamplingPercentage: jaegerSamplingPercentage,
		},
	}
}
