package manager

import (
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"
	"github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
)

type Option = func(*TransactionManager)

func WithClientOptions(option *client.Options) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.options = option
	}
}

func WithNewrelicSDK(sdk *newrelic.Application) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.newrelicSDK = sdk
	}
}

func WithLogger(logger *zap.Logger) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.logger = logger
	}
}

func WithPostgres(postgres *postgres.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.postgresClient = postgres
	}
}

func WithMongoConn(mongoConn *mongo.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.mongoClient = mongoConn
	}
}

func WithRetryPolicy(policy *Policy) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.retryPolicy = policy
	}
}

func WithRpcTimeout(timeout time.Duration) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.rpcTimeout = timeout
	}
}

func WithMetricsEnabled(enabled bool) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.metricsEnabled = enabled
	}
}
