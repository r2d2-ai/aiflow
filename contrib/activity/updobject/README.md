<!--
title: UpdObject
weight: 4616
-->

# UpdObject
This activity allows update an existing object's values.

## Installation

### aiflow CLI
```bash
aiflow install github.com/r2d2-ai/ai-box/contrib/activity/updobject
```

## Configuration

### Input:
| Name  | Type   | Description
|:---   | :---   | :---    
| object | object |  The object to update
| values | object |  The keys/properties and values to update or add

## Example

Update the passthru object's foo value:

```json
{
  "id": "updflowpassthru",
  "name": "Update Flow Passthru",
  "activity": {
    "ref": "github.com/r2d2-ai/ai-box/contrib/activity/updobject",
    "input": {
      "object": "=$flow.passthru",
      "values": {"foo": "bar"}
    }
  }
}
```