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

func (f *Client) AddActivity(ctx context.Context, feedID *string, activity *stream.Activity) error {
	// TODO: instrument this
	if activity == nil {
		return fmt.Errorf("invalid input argument. activity: %v", activity)
	}

	if feedID == nil {
		return fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return err
	}

	res, err := feed.AddActivity(ctx, *activity)
	if err != nil {
		return err
	}

	f.logger.Info(fmt.Sprintf("created follow request record via getstream. activityID: %s", res.Activity.ID))
	return nil
}

func (f *Client) DeleteActivity(ctx context.Context, feedID *string, activityForeignID *string) error {
	// TODO: instrument this
	if activityForeignID == nil {
		err := fmt.Errorf("invalid input argument. postID: %d", activityForeignID)
		return err
	}

	if feedID == nil {
		err := fmt.Errorf("invalid input argument. feedId: %v", feedID)
		return err
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return err
	}

	res, err := feed.RemoveActivityByForeignID(ctx, *activityForeignID)
	if err != nil {
		return err
	}

	f.logger.Info(fmt.Sprintf("activity removal attempt. status: %s", res.Removed))
	return nil
}

// AddActivityToManyFeeds adds an activity to many feeds
func (f *Client) AddActivityToManyFeeds(ctx context.Context, activity *stream.Activity, feeds ...stream.Feed) error {
	// TODO: instrument this
	if activity == nil {
		return fmt.Errorf("invalid activity object. activity cannot be nil")
	}

	if len(feeds) == 0 {
		return fmt.Errorf("invalid feed set. target feeds cannot be empty")
	}

	return f.client.AddToMany(ctx, *activity, feeds...)
}
