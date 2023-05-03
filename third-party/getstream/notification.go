// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package feeds

import (
	"context"

	"github.com/GetStream/stream-go2/v7"
)

func (f *Client) SendFollowRequestNotification(ctx context.Context, notificationFeedID string, params *FollowRequestActivity) error {
	notificationFeed, err := f.GetNotificationFeedFromFeedID(&notificationFeedID)
	if err != nil {
		return err
	}

	activity := &stream.Activity{
		Actor:     params.SourceActor,
		Verb:      params.ActionName,
		Object:    params.FollowRequestRecordID,
		ForeignID: params.ActivityForeignID,
		Time:      stream.Time{Time: params.Time},
	}

	if _, err := notificationFeed.AddActivity(ctx, *activity); err != nil {
		return err
	}

	return nil
}
