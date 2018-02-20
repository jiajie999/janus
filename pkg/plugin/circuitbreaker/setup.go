package circuitbreaker

import (
	"github.com/hellofresh/janus/pkg/plugin"
	"github.com/hellofresh/janus/pkg/proxy"
)

// Config represents the CORS configuration
type Config struct {
	Name                   string `json:"name"`
	Timeout                int    `json:"timeout"`
	MaxConcurrentRequests  int    `json:"max_concurrent_requests"`
	RequestVolumeThreshold int    `json:"request_volume_threshold"`
	SleepWindow            int    `json:"sleep_window"`
	ErrorPercentThreshold  int    `json:"error_percent_threshold"`
}

func init() {
	plugin.RegisterPlugin("circuit_breaker", plugin.Plugin{
		Action: setupCircuitBreaker,
	})
}

func setupCircuitBreaker(route *proxy.Route, rawConfig plugin.Config) error {
	var config Config

	err := plugin.Decode(rawConfig, &config)
	if err != nil {
		return err
	}

	route.Breaker = &Breaker{
		Name:                   config.Name,
		Timeout:                config.Timeout,
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		ErrorPercentThreshold:  config.ErrorPercentThreshold,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.SleepWindow,
	}

	return nil
}
