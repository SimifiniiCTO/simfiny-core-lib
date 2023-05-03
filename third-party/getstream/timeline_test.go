// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package feeds

import (
	"context"
	"github.com/SimifiniiCTO/simfiny-social-service/pkg/database/utils"
	"testing"
)

func TestClient_GetTimeline(t *testing.T) {
	userId := utils.GenerateRandomString(10)
	flatFeed := generateFlatFeed(t, userId)
	type args struct {
		ctx    context.Context
		feedID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - get a valid timeline",
			f:    testClient,
			args: args{
				ctx:    context.Background(),
				feedID: utils.StringP(flatFeed.ID()),
			},
			wantErr: false,
		},
		{
			name: "fail - get an invalid timeline",
			f:    testClient,
			args: args{
				ctx:    context.Background(),
				feedID: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.f.GetTimeline(tt.args.ctx, tt.args.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetTimeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetNotificationTimeline(t *testing.T) {
	userId := utils.GenerateRandomString(10)
	feed := generateNotificationFeed(t, userId)
	type args struct {
		ctx    context.Context
		feedID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "Pass - valid notification feed",
			f:    testClient,
			args: args{
				ctx:    context.Background(),
				feedID: utils.StringP(feed.ID()),
			},
			wantErr: false,
		},
		{
			name: "fail - invalid notification feed",
			f:    testClient,
			args: args{
				ctx:    context.Background(),
				feedID: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.f.GetNotificationTimeline(tt.args.ctx, tt.args.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNotificationTimeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
