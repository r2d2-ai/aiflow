<!--
title: Log
weight: 4615
-->

# Log
This activity allows you to write log messages.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/AIflow/common/activity/log
```

## Configuration

### Input:
| Name       | Type   | Description
|:---        | :---   | :---    
| message    | string | The message to log
| addDetails | bool   | Append contextual execution information to the log message
| usePrint   | bool   | Use Print to stdout instead of the built-in Logger. The built-in logger may not be compatible with some cloud environments. 

## Examples
The below example logs a message 'test message':

```json
{
  "id": "log_message",
  "name": "Log Message",
  "activity": {
    "ref": "github.com/r2d2-ai/AIflow/common/activity/log",
    "input": {
      "message": "test message",
      "addDetails": "false"
    }
  }
}
```