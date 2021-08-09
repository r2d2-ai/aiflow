
# Single Command Example
This example shows how to configure the CLI trigger to run as a single command.  It assumes
that there is one handler, which runs by default.

To build and run cli example and execute
```
aiflow create -f AIflow-single-cli.json
cd cli
aiflow build --shim cli
./bin/cli
```


## Configuration
```json
{
  "triggers": [
    {
      "id": "cli",
      "ref": "#cli",
      "name": "simple",
      "description": "Simple CLI Utility",
      "settings": {
        "singleCmd": true
      },
      "handlers": [
        {
          "name": "test1",
          "settings": {
            "usage": "[flags] [args]",
            "short": "test command",
            "long": "the test command",
            "flags": [
              "flag1||||the first value flag",
              "flag2||false||the first bool flag"
            ]
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:command1"
            },
            "input": {
              "flags": "=$.flags",
              "args": "=$.args"
            }
          }
        }
      ]
    }
  ]
 }
```