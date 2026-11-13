module github.com/tcero76/marketplace/redis/config

go 1.25.2

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/tcero76/marketplace/config/log v0.0.0-00010101000000-000000000000
)

require (
	github.com/bshuster-repo/logrus-logstash-hook v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace github.com/tcero76/marketplace/config/log => ../../config/log

replace github.com/tcero76/marketplace/config/metrics => ../../config/metrics
