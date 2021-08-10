module main

go 1.16

replace github.com/r2d2-ai/ai-box/core => ../../../../core
replace github.com/r2d2-ai/ai-box/flow => ../../../../flow
replace github.com/r2d2-ai/ai-box/contrib/activity/log => ../../../../contrib/activity/log
replace github.com/r2d2-ai/ai-box/cv/trigger/ipcam => ../../../../cv/trigger/ipcam

//require github.com/r2d2-ai/ai-box/core v1.0.0 // indirect
