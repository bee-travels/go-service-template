module {{.ServiceNamePill}}

go 1.15

require (
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/instana/go-sensor v1.29.0
	github.com/iris-contrib/middleware/cors v0.0.0-20210110101738-6d0a4d799b5d
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris/v12 v12.2.0-alpha2
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
)

replace github.com/russross/blackfriday/v2 => gopkg.in/russross/blackfriday.v2 v2.1.0
