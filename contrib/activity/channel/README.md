<!-- 
title: Channel
weight: 4603
-->

# Channel
This activity allows you to put a data on a named channel in the AIflow engine.  Channels are
essentially an internal communication channel in the engine.


## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/ai-box/contrib/activity/channel
```

## Configuration

### Input:
| Name    | Type   | Description
|:---     | :---   | :---    
| channel | string | The name of channel to use - **REQUIRED**
| data    | any    | The data to put on the channel

