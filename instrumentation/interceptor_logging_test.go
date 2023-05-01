package instrumentation

import (
	"reflect"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

func TestInterceptorLogger(t *testing.T) {
	type args struct {
		l *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want logging.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterceptorLogger(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterceptorLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
