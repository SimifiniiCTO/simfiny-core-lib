// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import (
	"reflect"
	"testing"
)

func Test_newRequestCountMetric(t *testing.T) {
	type args struct {
		serviceName *string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRequestCountMetric(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequestCountMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRequestLatencyMetric(t *testing.T) {
	type args struct {
		serviceName *string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRequestLatencyMetric(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequestLatencyMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newErrorCountMetric(t *testing.T) {
	type args struct {
		serviceName *string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newErrorCountMetric(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newErrorCountMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRequestStatusSummaryMetric(t *testing.T) {
	type args struct {
		serviceName *string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRequestStatusSummaryMetric(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequestStatusSummaryMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDbOperationCounter(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDbOperationCounter(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDbOperationCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDbOperationLatency(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDbOperationLatency(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDbOperationLatency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGrpcRequestCounter(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newGrpcRequestCounter(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGrpcRequestCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGrpcRequestLatency(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newGrpcRequestLatency(tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGrpcRequestLatency() = %v, want %v", got, tt.want)
			}
		})
	}
}
