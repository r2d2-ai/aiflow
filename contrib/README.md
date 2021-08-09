# AIflow Contrib

[![Build Status](https://travis-ci.org/r2d2-ai/ai-box/contrib.svg?branch=master)](https://travis-ci.org/r2d2-ai/ai-box/contrib.svg?branch=master)

Core collection of AIflow contributions.  This repository consists of activities, triggers and functions.

## Contributions

### Activities
* [actreply](activity/actreply): Action Reply
* [actreturn](activity/actreturn): Action Return
* [appdata](activity/actreturn): Application Shared Data
* [channel](activity/channel):  Internal Engine Message Publisher  
* [counter](activity/counter): Simple Counter 
* [error](activity/error): Throw Error
* [jsexec](activity/jsexec): Execute Javascript 
* [kafka](activity/kafka): Kafka Publisher
* [log](activity/log): Log Message
* [mapper](activity/mapper): Mapper
* [noop](activity/noop): No-Op 
* [rest](activity/rest): REST Invoker 
* [sqlquery](activity/sqlquery): Run SQL Query 

### Triggers
* [channel](trigger/channel): Internal Engine Message Listener
* [cli](trigger/cli): CLI
* [kafka](trigger/kafka): Kafka Subscriber
* [loadtester](trigger/loadtester): Basic Load Tester
* [rest](trigger/rest): REST
* [timer](trigger/timer): Timer
 
### Functions
* [coerce](function/coerce): Type Conversion
* [json](function/json): JSON
* [number](function/number): Numbers
* [string](function/string): Strings

## Installation

#### Install Activity
Example: install **log** activity

```bash
AIflow install github.com/r2d2-ai/ai-box/contrib/activity/log
```
#### Install Trigger
Example: install **rest** trigger

```bash
AIflow install github.com/r2d2-ai/ai-box/contrib/trigger/rest
```
#### Install Functions
Example: install **string** functions

```bash
AIflow install github.com/r2d2-ai/ai-box/contrib/function/string
```

## Contributing and support

### Contributing

New activities, triggers and functions are welcomed. If you would like to contribute, please following the [contribution guidelines](https://github.com/TIBCOSoftware/AIflow/blob/master/CONTRIBUTING.md). If you have any questions, issues, etc feel free to chat with us on [Gitter](https://gitter.im/r2d2-ai/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link).

## License
The contrib repository is licensed under a BSD-type license. See [LICENSE](LICENSE) for license text.

### Support
For Q&A you can post your questions on [Gitter](https://gitter.im/r2d2-ai/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)
