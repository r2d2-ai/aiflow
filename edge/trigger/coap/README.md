# CoAP
This trigger provides your AIflow application the ability to start a flow via CoAP.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/aiflow/edge-contrib/trigger/coap
```

## Configuration

### Settings:
| Name      | Type   | Description
| :---      | :---   | :---
| port    | string | 	The port number


### Handler Settings
| Name      | Type   | Description
| :---      | :---   | :---
| path | string | The path of the CoAP request
| method    | string | 	The Method to listen on

 
### Output: 

| Name    | Type   | Description
| :---    | :---   | :---
| queryParams | params | The query parameters
| payload | string | The payload of the CoAP Message
    

## Example

```json
{
  "id": "coap-trigger",
  "name": "Coap Trigger",
  "ref": "github.com/r2d2-ai/aiflow/edge-contrib/trigger/coap",
  "settings": {
      "port" : "5683"
  },
  "handlers": {
    "settings": {
    	"path": "/AIflow"
    
    }
  }
}
```