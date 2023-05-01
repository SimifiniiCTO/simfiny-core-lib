// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import (
	"reflect"
	"testing"

	"google.golang.org/grpc"
)

func TestClient_ErrorCountUnaryServerInterceptor(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		want grpc.UnaryServerInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ErrorCountUnaryServerInterceptor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ErrorCountUnaryServerInterceptor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ErrorCountStreamServerInterceptor(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		want grpc.StreamServerInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ErrorCountStreamServerInterceptor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ErrorCountStreamServerInterceptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
