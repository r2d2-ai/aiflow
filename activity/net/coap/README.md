<!--
title: CoAP
weight: 4607
-->

# CoAp
This activity allows you to send a CoAP message.

## Installation

### AIflow CLI
```bash
AIflow install github.com/r2d2-ai/aiflow/edge-contrib/activity/coap
```

## Configuration

### Settings:
| Name        | Type   | Description
| :---        | :---   | :---
| method      | string | The CoAP method to use (allowed values are GET, POST, PUT, DELETE)  - ***REQUIRED***   
| uri         | string | The CoAP resource URI - ***REQUIRED***
| messageType | string | The message type (allowed values are Confirmable, NonConfirmable, Acknowledgement, Reset),  *Confirmable* is the default 
| options     | params | The CoAP options to set     

### Input: 

| Name       | Type   | Description
| :---       | :---   | :---
| queryParams| params | The query params of the CoAP message    
| payload    | string | The payload of the CoAP message   
| messageId  | int    | ID used to detect duplicates and for optional reliability
 

### Output:

| Name       | Type   | Description
| :---       | :---   | :---
| response   | string | The response

## Example

```json
{
  "id": "coap-activity",
  "name": "Coap Activity",
  "description": "CoAP Get Example",
  "activity": {
    "ref": "github.com/r2d2-ai/aiflow/edge-contrib/activity/coap",
    "settings": {
      "method" : "GET",
      "uri": "coap://localhost:5683/AIflow"
    },
    "input" : {
        "payload" : "Hello from AIflow"
    }
  }
}
```