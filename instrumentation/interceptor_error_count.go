// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation // import "github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"

import (
	"context"
	"fmt"
	"time"

	rkgrpcctx "github.com/rookie-ninja/rk-grpc/interceptor/context"
	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	"google.golang.org/grpc"
)

// ErrorCountUnaryServerInterceptor Create new unary server interceptor to capture number of errors.
func (c *Client) ErrorCountUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = rkgrpcmid.WrapContextForServer(ctx)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))
		resp, err := handler(ctx, req)

		if err != nil {
			if c.Enabled {
				c.recordErrorMetric(c.baseMetrics.ErrorCountMetric, fmt.Sprintf("%s-%s", method, path), err.Error(), time.Now())
			}
		}

		return resp, err
	}
}

// ErrorCountStreamServerInterceptor Create new stream server interceptor to capture error count over time.
func (c *Client) ErrorCountStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		err := handler(srv, wrappedStream)

		if err != nil {
			if c.Enabled {
				c.recordErrorMetric(c.baseMetrics.ErrorCountMetric, fmt.Sprintf("%s-%s", method, path), err.Error(), time.Now())
			}
		}

		return err
	}
}
