<!--
title: No-Op
weight: 4615
-->

# No-Op
This activity is a simple No-Op that can be used for testing.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/AIflow/common/activity/noop
```

## Examples
Configuration of a No-Op activity

```json
{
  "id": "noop",
  "name": "NoOp",
  "activity": {
    "ref": "github.com/r2d2-ai/AIflow/common/activity/noop"
  }
}
```