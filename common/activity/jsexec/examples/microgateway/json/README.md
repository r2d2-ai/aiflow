# Gateway with Javascript
This recipe is a gateway which runs some javascript.

## Installation
* Install [Go](https://golang.org/)
* Install the AIflow [CLI](https://github.com/r2d2-ai/cli)

## Setup
```bash
git clone https://github.com/r2d2-ai/jsexec
cd jsexec/examples/microgateway/json
```

## Testing
Create the gateway:
```bash
AIflow create -f AIflow.json
cd MyProxy
AIflow build
```

Start the gateway:
```bash
bin/MyProxy
```

Run the following command:
```bash
curl http://localhost:9096/calculate"
```

You should see the following like response:
```json
{"sum":3}
```
