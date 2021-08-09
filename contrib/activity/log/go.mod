module github.com/r2d2-ai/ai-box/contrib/activity/log

retract v0.10.1-0.20210719201826-6001263f19e7

require (
	github.com/r2d2-ai/ai-box/core v1.0.0
	github.com/stretchr/testify v1.7.0
)

go 1.16

replace github.com/r2d2-ai/ai-box/core => ../../../core
