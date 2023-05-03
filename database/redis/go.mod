module github.com/SimifiniiCTO/simfiny-core-lib/database/redis

go 1.20

replace github.com/SimifiniiCTO/simfiny-core-lib/instrumentation v0.0.0-00010101000000-000000000000 => ../../instrumentation

require (
	github.com/SimifiniiCTO/simfiny-core-lib/instrumentation v1.0.0
	github.com/alicebob/miniredis/v2 v2.30.2
	github.com/gomodule/redigo v1.8.9
	github.com/stretchr/testify v1.8.2
	go.uber.org/zap v1.24.0
)

require (
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/newrelic/go-agent v3.21.1+incompatible // indirect
	github.com/newrelic/go-agent/v3 v3.21.1 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrzap v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/yuin/gopher-lua v1.1.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
