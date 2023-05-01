// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import (
	"testing"
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
)

func TestClient_recordDtxMetric(t *testing.T) {
	type args struct {
		m          *Metric
		op         string
		dest       string
		statusCode code.Code
		start      time.Duration
	}
	tests := []struct {
		name string
		c    *Client
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.recordDtxMetric(tt.args.m, tt.args.op, tt.args.dest, tt.args.statusCode, tt.args.start)
		})
	}
}

func TestClient_recordErrorMetric(t *testing.T) {
	type args struct {
		m               *Metric
		op              string
		msg             string
		timeOfOccurence time.Time
	}
	tests := []struct {
		name string
		c    *Client
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.recordErrorMetric(tt.args.m, tt.args.op, tt.args.msg, tt.args.timeOfOccurence)
		})
	}
}
