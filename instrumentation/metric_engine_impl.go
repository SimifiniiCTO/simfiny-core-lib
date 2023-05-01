// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation // import "github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"

import (
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
)

// RecordStandardDtxMetrics records standard distributed transaction metrics.
func (c *Client) RecordStandardDtxMetrics(op, dest string, status code.Code, start time.Time) {
	metricName := format(&MetricName{
		ServiceName:     c.GetServiceName(),
		OperationName:   op,
		IsDistributedTx: true,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Latency)

	c.recordDtxMetric(c.baseMetrics.RequestCountMetric, *metricName, dest, status, time.Since(start))
}
