// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"fmt"
	"strings"

	"github.com/GetStream/stream-go2/v7"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"go.uber.org/zap"
)

type FeedType string

const (
	PersonalFeed     FeedType = "personal"
	TimelineFeed     FeedType = "timeline"
	NotificationFeed FeedType = "notification"
)

func (f FeedType) String() string {
	return string(f)
}

// This is defining a struct type called `Client` which will be used to create instances of the
// GetStream client. The struct contains fields for the GetStream API key, secret, a pointer to a
// `stream.Client` instance, a logger, and an instrumentation client.
type Client struct {
	// `getstreamKey` is a field in the `Client` struct that stores the GetStream API key as a string. It
	// is used to authenticate requests to the GetStream API.
	getstreamKey string
	// `getstreamSecret` is a field in the `Client` struct that stores the GetStream API secret as a
	// string. It is used to authenticate requests to the GetStream API along with the API key.
	getstreamSecret string
	// `client                *stream.Client` is a field in the `Client` struct that stores a pointer to a
	// `stream.Client` instance. This instance is used to make requests to the GetStream API.
	client *stream.Client
	// `logger *zap.Logger` is a field in the `Client` struct that stores a pointer to a `zap.Logger`
	// instance. This instance is used for logging messages related to the GetStream client. `zap` is a
	// popular logging library in the Go programming language.
	logger *zap.Logger
	// `instrumentationClient *instrumentation.Client` is a field in the `Client` struct that stores a
	// pointer to an `instrumentation.Client` instance. This instance is used for instrumenting the
	// GetStream client, which means collecting and analyzing data related to the performance and behavior
	// of the client. This can help identify issues and optimize the client's performance.
	instrumentationClient *instrumentation.Client
}

// The function creates a new client with optional configuration options.
func New(opts ...Option) (*Client, error) {
	c := &Client{}
	c.ApplyOptions(opts...)
	if err := c.Validate(); err != nil {
		return nil, err
	}

	streamClient, err := stream.New(c.getstreamKey, c.getstreamSecret)
	if err != nil {
		return nil, err
	}

	c.client = streamClient
	return c, nil
}

// `func (f *Client) GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error)` is a method of
// the `Client` struct that takes a pointer to a string as input and returns a pointer to a
// `stream.FlatFeed` object and an error. It is used to retrieve a `FlatFeed` object from the GetStream
// API based on a given feed ID. The method first validates the feed ID using the `validateFeedID`
// method, and then splits the feed ID into its constituent parts (feed type and feed slug) using the
// `strings.Split` function. Finally, it calls the `FlatFeed` method of the `stream.Client` instance
// stored in the `Client` struct, passing in the feed type and feed slug as arguments, and returns the
// resulting `FlatFeed` object and any error that occurred during the API call.
func (f *Client) GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.FlatFeed(details[0], details[1])
}

// `func (f *Client) GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error)`
// is a method of the `Client` struct that takes a pointer to a string as input and returns a pointer
// to a `stream.NotificationFeed` object and an error. It is used to retrieve a `NotificationFeed`
// object from the GetStream API based on a given feed ID. The method first validates the feed ID using
// the `validateFeedID` method, and then splits the feed ID into its constituent parts (feed type and
// feed slug) using the `strings.Split` function. Finally, it calls the `NotificationFeed` method of
// the `stream.Client` instance stored in the `Client` struct, passing in the feed type and feed slug
// as arguments, and returns the resulting `NotificationFeed` object and any error that occurred during
// the API call.
func (f *Client) GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.NotificationFeed(details[0], details[1])
}

// `func (f *Client) validateFeedID(feedID *string) error` is a method of the `Client` struct that
// takes a pointer to a string as input and returns an error. It is used to validate the feed ID passed
// as input to other methods of the `Client` struct. The method checks if the `Client` instance and the
// feed ID are not nil, and returns an error if either of them is nil. If both are not nil, the method
// returns nil, indicating that the feed ID is valid.
func (f *Client) validateFeedID(feedID *string) error {
	if f.client == nil || feedID == nil {
		return fmt.Errorf("invalid input argument. client: %v, feedId :%v", f.client, feedID)
	}
	return nil
}

// `func (c *Client) CreateFlatFeed(feedType, feedSlug string) (*stream.FlatFeed, error)` is a method
// of the `Client` struct that creates a new `FlatFeed` object in the GetStream API based on the given
// feed type and feed slug. It takes the feed type and feed slug as input arguments and returns a
// pointer to the newly created `FlatFeed` object and any error that occurred during the API call. The
// method first checks if the input arguments are valid, and then calls the `FlatFeed` method of the
// `stream.Client` instance stored in the `Client` struct, passing in the feed type and feed slug as
// arguments.
func (c *Client) CreateFlatFeed(feedType FeedType, feedSlug string) (*stream.FlatFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.FlatFeed(feedType.String(), feedSlug)
}

// `func (c *Client) CreateAggregateFeed(feedType, feedSlug string) (*stream.AggregatedFeed, error)` is
// a method of the `Client` struct that creates a new `AggregatedFeed` object in the GetStream API
// based on the given feed type and feed slug. It takes the feed type and feed slug as input arguments
// and returns a pointer to the newly created `AggregatedFeed` object and any error that occurred
// during the API call. The method first checks if the input arguments are valid, and then calls the
// `AggregatedFeed` method of the `stream.Client` instance stored in the `Client` struct, passing in
// the feed type and feed slug as arguments.
func (c *Client) CreateAggregateFeed(feedType FeedType, feedSlug string) (*stream.AggregatedFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.AggregatedFeed(feedType.String(), feedSlug)
}

// `func (c *Client) CreateNotificationFeed(feedType, feedSlug string) (*stream.NotificationFeed,
// error)` is a method of the `Client` struct that creates a new `NotificationFeed` object in the
// GetStream API based on the given feed type and feed slug. It takes the feed type and feed slug as
// input arguments and returns a pointer to the newly created `NotificationFeed` object and any error
// that occurred during the API call. The method first checks if the input arguments are valid, and
// then calls the `NotificationFeed` method of the `stream.Client` instance stored in the `Client`
// struct, passing in the feed type and feed slug as arguments.
func (c *Client) CreateNotificationFeed(feedType FeedType, feedSlug string) (*stream.NotificationFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.NotificationFeed(feedType.String(), feedSlug)
}
