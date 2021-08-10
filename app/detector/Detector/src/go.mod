module main

go 1.16

require (
	github.com/r2d2-ai/ai-box/contrib/activity/log v1.0.0
	github.com/r2d2-ai/ai-box/core v1.0.0
	github.com/r2d2-ai/ai-box/cv/trigger/ipcam v1.0.0
	github.com/r2d2-ai/ai-box/flow v1.0.0
)

replace github.com/r2d2-ai/ai-box/core => ../../../../core

replace github.com/r2d2-ai/ai-box/flow => ../../../../flow

replace github.com/r2d2-ai/ai-box/contrib/activity/log => ../../../../contrib/activity/log

replace github.com/r2d2-ai/ai-box/cv/trigger/ipcam => ../../../../cv/trigger/ipcam
