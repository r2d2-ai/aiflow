// Do not change this file, it has been generated using AIflow-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded AIflow app descriptor file
const AIflowJSON string = `{
  "name": "Detector",
  "type": "AIflow:app",
  "version": "0.0.1",
  "appModel": "1.0.0",
  "description": "",
  "imports": [
    "github.com/r2d2-ai/ai-box/contrib/activity/log",
    "github.com/r2d2-ai/ai-box/cv/trigger/ipcam",
    "github.com/r2d2-ai/ai-box/flow"
  ],
  "triggers": [
    {
      "id": "ai_box_ipcam",
      "ref": "#ipcam",
      "name": "IP Cam Feed",
      "description": "IP Cam Feed",
      "handlers": [
        {
          "settings": {
            "protocol": "RSTP",
            "host": "192.168.50.195",
            "user": "admin",
            "password": "P$rolaMeaCAM",
            "videoUri": "11",
            "groupId": "group-1",
            "cameraId": "cam-1"
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:detect"
            },
            "input": {
              "groupId": "=$.groupId",
              "cameraId": "=$.cameraId",
              "image": "=$.image"
            }
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:detect",
      "data": {
        "name": "Detect",
        "metadata": {
          "input": [
            {
              "name": "cameraId",
              "type": "string"
            },
            {
              "name": "groupId",
              "type": "string"
            },
            {
              "name": "image",
              "type": "any"
            }
          ]
        },
        "tasks": [
          {
            "id": "log_2",
            "name": "Log",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$flow.cameraId",
                "addDetails": false,
                "usePrint": false
              }
            }
          }
        ]
      }
    }
  ]
}
`
const engineJSON string = ``

func init() {
	cfgJson = AIflowJSON
	cfgEngine = engineJSON
}
