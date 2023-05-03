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

func (f *Client) FollowFeed(ctx context.Context, sourceFeedID, targetFeedID string, opts []stream.FollowFeedOption) error {
	if sourceFeedID == "" || targetFeedID == "" {
		return fmt.Errorf("invalid input argument. sourceFeedID: %s, targetFeedID: %s", sourceFeedID, targetFeedID)
	}

	sourceFeed, err := f.GetFlatFeedFromFeedID(&sourceFeedID)
	if err != nil {
		return err
	}

	targetFeed, err := f.GetFlatFeedFromFeedID(&targetFeedID)
	if err != nil {
		return err
	}

	if _, err = sourceFeed.Follow(ctx, targetFeed, opts...); err != nil {
		return err
	}

	f.logger.Info(fmt.Sprintf("created follow record via getstream. target: %s, source: %s", targetFeedID, sourceFeedID))
	return nil
}
