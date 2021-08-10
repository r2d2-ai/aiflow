# RunAction
This activity allows you to run AIflow actions.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/AIflow/common/activity/runaction
```

## Configuration

### Settings
| Name          | Type   | Description
|:---           | :---   | :---    
| actionRef     | string | The 'ref' to the action to be run
| actionSettings| object | The settings of the action

### Input
The inputs for this activity should be the inputs for the action you are running

### Output
| Name          | Type   | Description
|:---           | :---   | :---    
| output        | object | The output of the action.


## Examples
The below example logs a message 'test message':

```json
{
    "id": "cmlact",
    "ref": "github.com/r2d2-ai/AIflow/common/activity/runaction",
    "settings": {
        "actionRef": "github.com/r2d2-ai/catalystml-AIflow/action",
        "actionSettings": { "catalystMlURI" : "file://cml.json" }
    },
    "input": {
        "dataIn": "=$.result"
    }
}          
```
