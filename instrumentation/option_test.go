package instrumentation

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestWithServiceName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithServiceName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithServiceVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithServiceVersion(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithServiceVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithServiceEnvironment(t *testing.T) {
	type args struct {
		environment string
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithServiceEnvironment(tt.args.environment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithServiceEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEnabled(t *testing.T) {
	type args struct {
		enabled bool
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEnabled(tt.args.enabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithNewrelicKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithNewrelicKey(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithNewrelicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want ServiceTelemetryOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Validate(t *testing.T) {
	tests := []struct {
		name    string
		tr      *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
