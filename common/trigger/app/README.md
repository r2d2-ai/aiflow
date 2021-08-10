<!--
title: APP
weight: 4706
-->

# App Trigger
This trigger provides your AIflow application the ability to start an action on the Lifecycle events of the 
application.  The handler associated with the *STARTUP* event gets invoked before all other triggers have 
been started. The handler associated with the *SHUTDOWN* event gets invoked after all other triggers have 
been stopped.  

## Installation

```bash
AIflow install github.com/r2d2-ai/aiflow/common/trigger/app
```

## Configuration

### Handler Settings:
| Name     | Type   | Description
|:---      | :---   | :---          
| lifecycle   | string | The lifecycle event (ie. STARTUP,SHUTDOWN) - **REQUIRED**

## Example

Configure the trigger in order to have a special startup and cleanup action for your application.

```json
{
  "triggers": [
    {
      "id": "AIflow-app",
      "ref": "github.com/r2d2-ai/aiflow/common/trigger/app",
      "handlers": [
        {
          "settings": {
            "lifecycle": "STARTUP"
          },
          "action": {
            "ref": "github.com/r2d2-ai/aiflow/flow",
            "settings": {
              "flowURI": "res://flow:app_startup"
            }
          }
        },
        {
          "settings": {
            "lifecycle": "SHUTDOWN"
          },
          "action": {
            "ref": "github.com/r2d2-ai/aiflow/flow",
            "settings": {
              "flowURI": "res://flow:app_cleanup"
            }
          }
        }
      ]
    }
  ]
}
```
