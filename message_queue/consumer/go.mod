module github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer

go 1.20

replace github.com/SimifiniiCTO/simfiny-core-lib/instrumentation v1.0.1 => ../../instrumentation

replace github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client v1.0.1 => ../client

require (
	github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client v1.0.1
	github.com/newrelic/go-agent/v3 v3.21.1
	go.uber.org/zap v1.24.0
)

require (
	github.com/aws/aws-sdk-go v1.44.253 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.2.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
