<!--
title: Channel
weight: 4706
-->

# Channel Trigger
This trigger provides your AIflow application the ability to start an action via a named engine channel

## Installation

```bash
AIflow install github.com/r2d2-ai/aiflow/common/trigger/channel
```

## Configuration    

### Handler Settings:
| Name    | Type   | Description
|:---     | :---   | :---     
| channel | string | The internal engine channel - **REQUIRED**

#### Output:
| Name | Type | Description
|:---  | :--- | :---     
| data | any  | The data pulled from the channel


## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the Channel Trigger.

### Run Flow
Configure the Trigger to handle an event received on the 'test' channel

```json
{
  "triggers": [
    {
      "id": "AIflow-channel",
      "ref": "github.com/r2d2-ai/aiflow/common/trigger/channel",
      "handlers": [
        {
          "settings": {
            "channel": "test"
          },
          "action": {
            "ref": "github.com/r2d2-ai/aiflow/action/flow",
            "settings": {
                "flowURI": "res://flow:testflow"
            }       
          }
        }
      ]
    }
  ]
}
```

Note: a channel must be defined in the app in order to use it in a trigger or activity.  A channel is specified by "\<name\>:\<buffer size\>"
```
"channels": [
  "test:5"
]
```