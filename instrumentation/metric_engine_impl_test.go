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

func TestClient_RecordStandardDtxMetrics(t *testing.T) {
	type args struct {
		op     string
		dest   string
		status code.Code
		start  time.Time
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
			tt.c.RecordStandardDtxMetrics(tt.args.op, tt.args.dest, tt.args.status, tt.args.start)
		})
	}
}
