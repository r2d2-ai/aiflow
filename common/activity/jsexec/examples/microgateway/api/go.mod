module github.com/r2d2-ai/aiflow/common/activity/jsexec

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/r2d2-ai/aiflow/common/trigger/rest v0.0.0-00010101000000-000000000000
	github.com/r2d2-ai/aiflow/core v0.9.4-hf.1
	github.com/r2d2-ai/microgateway v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

go 1.16

replace github.com/r2d2-ai/aiflow/core => ../../../../../../core

replace github.com/r2d2-ai/aiflow/common/trigger/rest => ../../../../../trigger/rest

replace github.com/r2d2-ai/aiflow/common/trigger/channel => ../../../../../trigger/channel

replace github.com/r2d2-ai/aiflow/common/activity/channel => ../../../../channel

replace github.com/r2d2-ai/aiflow/common/activity/counter => ../../../../counter

replace github.com/r2d2-ai/aiflow/common/activity/log => ../../../../log

replace github.com/r2d2-ai/aiflow/common/activity/rest => ../../../../rest

replace github.com/r2d2-ai/aiflow/common/function/coerce => ../../../../../function/coerce

replace github.com/r2d2-ai/aiflow/common/function/json => ../../../../../function/json

replace github.com/r2d2-ai/aiflow/common/function/number => ../../../../../function/number

replace github.com/r2d2-ai/aiflow/common/function/string => ../../../../../function/string

replace github.com/r2d2-ai/microgateway => ../../../../../../microgateway
