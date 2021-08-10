<!--
title: Mapper
weight: 4616
-->

# Mapper
This activity allows you to map values to the working attribute set of an action.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/aiflow/common/activity/mapper
```

## Configuration

### Settings:
| Name     | Type   | Description
|:---      | :---   | :---     
| mappings | object | Set of mappings to execute

## Example
The below example allows you to configure the activity to map the output 'value' of activity 'myActivity' to FlowAttr1

```json
{
  "id": "mapper",
  "name": "Mapper",
  "activity": {
    "ref": "github.com/r2d2-ai/aiflow/common/activity/mapper",
    "input": {
      "mappings": 
        {
          "FlowAttr1": "=$activity[myActivity].value"
        }
      ]
    }
  }
}
```
