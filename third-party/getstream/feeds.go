// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package feeds

import (
	"fmt"
	"strings"

	"github.com/GetStream/stream-go2/v7"
	"go.uber.org/zap"
)

type Client struct {
	client *stream.Client
	logger *zap.Logger
}

func New(client *stream.Client, logger *zap.Logger) (*Client, error) {
	if client == nil || logger == nil {
		return nil, fmt.Errorf("invalid input argument. client: %v, logger: %v", client, logger)
	}

	return &Client{
		client: client,
		logger: logger,
	}, nil
}

// GetFlatFeedFromFeedID gets a flat feed object from a feed ID
func (f *Client) GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.FlatFeed(details[0], details[1])
}

// GetNotificationFeedFromFeedID gets a notification feed object (aggregated) from a feed ID
func (f *Client) GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.NotificationFeed(details[0], details[1])
}

// validateFeedID validates a feed id
func (f *Client) validateFeedID(feedID *string) error {
	if f.client == nil || feedID == nil {
		return fmt.Errorf("invalid input argument. client: %v, feedId :%v", f.client, feedID)
	}
	return nil
}

func (c *Client) CreateFlatFeed(feedType, feedSlug string) (*stream.FlatFeed, error) {
	// TODO: instrument this
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.FlatFeed(feedType, feedSlug)
}

func (c *Client) CreateAggregateFeed(feedType, feedSlug string) (*stream.AggregatedFeed, error) {
	// TODO: instrument this
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.AggregatedFeed(feedType, feedSlug)
}

func (c *Client) CreateNotificationFeed(feedType, feedSlug string) (*stream.NotificationFeed, error) {
	// TODO: instrument this
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.NotificationFeed(feedType, feedSlug)
}
