<p align="center">
  <img src ="https://raw.githubusercontent.com/TIBCOSoftware/AIflow/master/images/AIflow-ecosystem_Flows.png" />
</p>

<p align="center" >
  <b>Flows is a simple process flow action for the Project aiflow Ecosystem</b>
</p>

<p align="center">
  <img src="https://travis-ci.org/r2d2-ai/flow.svg?branch=master"/>
  <img src="https://img.shields.io/badge/dependencies-up%20to%20date-green.svg"/>
  <img src="https://img.shields.io/badge/license-BSD%20style-blue.svg"/>
  <a href="https://gitter.im/r2d2-ai/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link"><img src="https://badges.gitter.im/Join%20Chat.svg"/></a>
</p>

# aiflow Flow

aiflow Flow is one of the primary actions of the Project aiflow Ecosystem.  It is a process flow engine that can be used to connect activities together to create complex application logic. This proces engine can be used to create code-free applications or for application integration and orchestration projects.


## Getting Started

We’ve made building process flows as easy as possible. Develop your flows using:

- A simple, clean JSON-based DSL
- Golang API

See the sample below of flow that logs its inputs (for brevity, the triggers and metadata of the resource has been omitted). Also don’t forget to check out the [examples](https://github.com/r2d2-ai/ai-box/flow/tree/master/examples) in the repo.

```json
{
    "tasks": [
      {
        "id": "log_1",
        "name": "Log 1",
        "activity": {
          "ref": "#log",
          "input": {
            "message": "=$flow.in"
          }
        }
      },
      {
        "id": "log_2",
        "name": "Log 2",
        "activity": {
          "ref": "#log",
          "input": {
            "message": "second log message"
          }
        }
      }
    ],
    "links": [
      {
        "from":"log_1",
        "to":"log_2"
      }
    ]
}
```

## Try out the example

Firstly you should install the install the [aiflow CLI](https://github.com/r2d2-ai/cli).
 
Next you should download our aggregation example [log-aiflow.json](https://github.com/r2d2-ai/ai-box/flow/blob/master/examples/log-aiflow.json).

We'll create a our application using the example file, we'll call it myApp

```bash
$ aiflow create -f log-aiflow.json myApp
```

Now, build it...

```bash
$ cd myApp/
$ aiflow build
```

## Activities

aiflow Flows also provides some activities to assist in complex flow creation.

* [Subflow](activity/subflow/README.md) : This activity allows you start another flow within the existing flow.

## Additional Documentation

To learn more about the model you should checkout the [Flow Model](docs/model.md) doc.

## License 
aiflow source code in [this](https://github.com/r2d2-ai/strem) repository is under a BSD-style license, refer to [LICENSE](https://github.com/r2d2-ai/ai-box/flow/blob/master/LICENSE)
