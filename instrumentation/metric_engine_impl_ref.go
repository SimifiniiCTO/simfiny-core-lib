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

func (c *Client) recordDtxMetric(m *Metric, op, dest string, statusCode code.Code, start time.Duration) {
	if !c.Enabled {
		return
	}

	c.client.RecordCustomEvent(m.MetricName, map[string]interface{}{
		"service.source":      m.ServiceName,
		"service.operation":   op,
		"service.destination": dest,
		"metric.help":         m.Help,
		"namespace":           m.Namespace,
		"subsystem":           m.Subsystem,
		"duration":            start,
		"status.code":         statusCode,
	})
}

func (c *Client) recordErrorMetric(m *Metric, op, msg string, timeOfOccurence time.Time) {
	if !c.Enabled {
		return
	}

	c.client.RecordCustomEvent(m.MetricName, map[string]interface{}{
		"service.source":    m.ServiceName,
		"service.operation": op,
		"error.message":     msg,
		"metric.help":       m.Help,
		"metric.namespace":  m.Namespace,
		"metric.subsystem":  m.Subsystem,
		"metric.occurence":  timeOfOccurence.String(),
	})
}
