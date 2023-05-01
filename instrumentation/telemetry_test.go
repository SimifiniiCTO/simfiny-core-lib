package instrumentation

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
)

func TestNewServiceTelemetry(t *testing.T) {
	type args struct {
		opts []ServiceTelemetryOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServiceTelemetry(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServiceTelemetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceTelemetry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_IsEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEnabled(); got != tt.want {
				t.Errorf("Client.IsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetServiceEnvironment(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetServiceEnvironment(); got != tt.want {
				t.Errorf("Client.GetServiceEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetServiceName(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetServiceName(); got != tt.want {
				t.Errorf("Client.GetServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetServiceVersion(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetServiceVersion(); got != tt.want {
				t.Errorf("Client.GetServiceVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTraceFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetTraceFromContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetTraceFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTraceFromRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetTraceFromRequest(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetTraceFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTracingEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetTracingEnabled(); got != tt.want {
				t.Errorf("Client.GetTracingEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewChildSpan(t *testing.T) {
	type args struct {
		name   string
		parent newrelic.Transaction
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.Segment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NewChildSpan(tt.args.name, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewChildSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartTransaction(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartTransaction(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_WithContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		trace newrelic.Transaction
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WithContext(tt.args.ctx, tt.args.trace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_WithRequest(t *testing.T) {
	type args struct {
		r     *http.Request
		trace newrelic.Transaction
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WithRequest(tt.args.r, tt.args.trace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WithRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartExternalSegment(t *testing.T) {
	type args struct {
		txn *newrelic.Transaction
		req *http.Request
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.ExternalSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartExternalSegment(tt.args.txn, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartExternalSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRoundtripper(t *testing.T) {
	type args struct {
		txn *newrelic.Transaction
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *http.RoundTripper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NewRoundtripper(tt.args.txn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRoundtripper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartDatastoreSegment(t *testing.T) {
	type args struct {
		txn       *newrelic.Transaction
		operation string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartDatastoreSegment(tt.args.txn, tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartDatastoreSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartRedisDatastoreSegment(t *testing.T) {
	type args struct {
		txn       *newrelic.Transaction
		operation string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartRedisDatastoreSegment(tt.args.txn, tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartRedisDatastoreSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartNosqlDatastoreSegment(t *testing.T) {
	type args struct {
		txn            *newrelic.Transaction
		operation      string
		collectionName string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartNosqlDatastoreSegment(tt.args.txn, tt.args.operation, tt.args.collectionName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartNosqlDatastoreSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartMessageQueueSegment(t *testing.T) {
	type args struct {
		txn       *newrelic.Transaction
		queueName string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.MessageProducerSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartMessageQueueSegment(tt.args.txn, tt.args.queueName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartMessageQueueSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StartSegment(t *testing.T) {
	type args struct {
		txn  *newrelic.Transaction
		name string
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want *newrelic.Segment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartSegment(tt.args.txn, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUnaryServerInterceptors(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want []grpc.UnaryServerInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetUnaryServerInterceptors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetUnaryServerInterceptors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStreamServerInterceptors(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want []grpc.StreamServerInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetStreamServerInterceptors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStreamServerInterceptors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUnaryClientInterceptors(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want []grpc.UnaryClientInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetUnaryClientInterceptors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetUnaryClientInterceptors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStreamClientInterceptors(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want []grpc.StreamClientInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetStreamClientInterceptors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStreamClientInterceptors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewMuxRouter(t *testing.T) {
	tests := []struct {
		name string
		s    *Client
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NewMuxRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewMuxRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ConfigureNewrelicClient(t *testing.T) {
	tests := []struct {
		name    string
		s       *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.ConfigureNewrelicClient(); (err != nil) != tt.wantErr {
				t.Errorf("Client.ConfigureNewrelicClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
