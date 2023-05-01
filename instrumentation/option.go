package instrumentation

import (
	"fmt"

	"go.uber.org/zap"
)

// ServiceTelemetryOption is a function that configures a Client
type ServiceTelemetryOption func(*Client)

// WithServiceName configures the service name
func WithServiceName(name string) ServiceTelemetryOption {
	return func(t *Client) {
		t.ServiceName = name
	}
}

// WithServiceVersion configures the service version
func WithServiceVersion(version string) ServiceTelemetryOption {
	return func(t *Client) {
		t.ServiceVersion = version
	}
}

// WithServiceEnvironment configures the service environment
func WithServiceEnvironment(environment string) ServiceTelemetryOption {
	return func(t *Client) {
		t.ServiceEnvironment = environment
	}
}

// WithEnabled configures wether or not instrumentation is enabled
func WithEnabled(enabled bool) ServiceTelemetryOption {
	return func(t *Client) {
		t.Enabled = enabled
	}
}

// WithNewrelicKey configures the newrelic key
func WithNewrelicKey(key string) ServiceTelemetryOption {
	return func(t *Client) {
		t.NewrelicKey = key
	}
}

// WithLogger configures the logger
func WithLogger(logger *zap.Logger) ServiceTelemetryOption {
	return func(t *Client) {
		t.Logger = logger
	}
}

func (t *Client) Validate() error {
	if t.ServiceName == "" {
		return fmt.Errorf("service name not set")
	}

	if t.ServiceVersion == "" {
		return fmt.Errorf("service version not set")
	}

	if t.ServiceEnvironment == "" {
		return fmt.Errorf("service environment not set")
	}

	if t.NewrelicKey == "" {
		return fmt.Errorf("newrelic key not set")
	}

	if t.Logger == nil {
		return fmt.Errorf("logger not set")
	}

	if t.client == nil {
		return fmt.Errorf("client not set")
	}

	return nil
}
