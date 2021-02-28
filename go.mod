module github.com/fengjijiao/dingding-push-restful-api

go 1.15

replace github.com/fengjijiao/dingding-push-restful-api/pkg/logio => ./pkg/logio

replace github.com/fengjijiao/dingding-push-restful-api/pkg/conf => ./pkg/conf

replace github.com/fengjijiao/dingding-push-restful-api/pkg/sqlhandler => ./pkg/sqlhandler

replace github.com/fengjijiao/dingding-push-restful-api/pkg/httphandler => ./pkg/httphandler

require (
	github.com/jasonlvhit/gocron v0.0.1 // indirect
	go.uber.org/zap v1.16.0 // indirect
)
