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

func Test_newServiceBaseMetrics(t *testing.T) {
	type args struct {
		serviceName *string
	}
	tests := []struct {
		name    string
		args    args
		want    *ServiceBaseMetrics
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newServiceBaseMetrics(tt.args.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("newServiceBaseMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServiceBaseMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
