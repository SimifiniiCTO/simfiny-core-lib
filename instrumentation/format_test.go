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

func Test_format(t *testing.T) {
	type args struct {
		m          *MetricName
		metricType MetricType
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.args.m, tt.args.metricType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appenSuffix(t *testing.T) {
	type args struct {
		metricName string
		metricType MetricType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appenSuffix(tt.args.metricName, tt.args.metricType); got != tt.want {
				t.Errorf("appenSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
