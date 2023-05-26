module alvinr.ca/go-framework/examples

require (
	alvinr.ca/go-framework/base v0.0.0
	alvinr.ca/go-framework/http v0.0.0
	alvinr.ca/go-framework/lang v0.0.0
	alvinr.ca/go-framework/log v0.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

replace alvinr.ca/go-framework/base v0.0.0 => ../base

replace alvinr.ca/go-framework/lang v0.0.0 => ../lang

replace alvinr.ca/go-framework/log v0.0.0 => ../log

replace alvinr.ca/go-framework/http v0.0.0 => ../http

go 1.19
