// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package feeds

import (
	"context"
	"fmt"

	"github.com/GetStream/stream-go2/v7"
)

func (f *Client) GetTimeline(ctx context.Context, feedID *string) ([]stream.Activity, error) {
	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	resp, err := feed.GetActivities(ctx, stream.WithActivitiesLimit(100))
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func (f *Client) GetNotificationTimeline(ctx context.Context, feedID *string) ([]stream.NotificationFeedResult, error) {
	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetNotificationFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	resp, err := feed.GetActivities(ctx, stream.WithActivitiesLimit(100))
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}
