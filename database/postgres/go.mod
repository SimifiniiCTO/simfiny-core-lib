module github.com/SimifiniiCTO/simfiny-core-lib/database/postgres

go 1.20

replace github.com/SimifiniiCTO/simfiny-core-lib/instrumentation v0.0.0-00010101000000-000000000000 => ../../instrumentation

require (
	github.com/SimifiniiCTO/simfiny-core-lib/instrumentation v0.0.0-00010101000000-000000000000
	github.com/giantswarm/retry-go v0.0.0-20151203102909-d78cea247d5e
	github.com/jinzhu/gorm v1.9.16
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.25.0
)

require github.com/mattn/go-sqlite3 v1.14.15 // indirect

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/juju/errgo v0.0.0-20140925100237-08cceb5d0b53 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/newrelic/go-agent v3.21.1+incompatible // indirect
	github.com/newrelic/go-agent/v3 v3.21.1 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrzap v1.0.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gorm.io/driver/sqlite v1.5.0
)
