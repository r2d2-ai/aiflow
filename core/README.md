 
<p align="center">
  <img src ="https://raw.githubusercontent.com/TIBCOSoftware/AIflow/master/images/projectAIflow.png" />
</p>

<p align="center" >
  <b>Serverless functions and edge microservices made painless</b>
</p>

<p align="center">
  <img src="https://travis-ci.org/r2d2-ai/core.svg?branch=master"/>
  <img src="https://img.shields.io/badge/dependencies-up%20to%20date-green.svg"/>
  <img src="https://img.shields.io/badge/license-BSD%20style-blue.svg"/>
  <a href="https://gitter.im/r2d2-ai/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link"><img src="https://badges.gitter.im/Join%20Chat.svg"/></a>
</p>

<p align="center">
  <a href="#getting-started">Getting Started</a> | <a href="#documentation">Documentation</a> | <a href="#contributing">Contributing</a> | <a href="#license">License</a>
</p>

<br/>

Project aiflow is an open source framework to simplify building efficient & modern serverless functions and edge microservices and _this_ repository is the core library used to create and extend those **aiflow Applications**. 

# aiflow Core
aiflow Core is the core aiflow library which contains the apis to create and extend aiflow applications.

## Getting started
If you want to get started with [Project AIflow](AIflow.io), you should install the install the [aiflow CLI](https://github.com/r2d2-ai/ai-box/cli).  You can find details there on creating a quick sample application.  You also might want to check out the [getting started](https://tibcosoftware.github.io/AIflow/getting-started/) guide in our docs or check out the [Labs](https://tibcosoftware.github.io/AIflow/labs/) section in our docs for in depth tutorials.

## Documentation
Here is some documentation to help you get started understanding some of the fundamentals of the aiflow Core library. 

* [Model](docs/model.md): The aiflow application model
* [Data Types](docs/datatypes.md): The aiflow data types
* [Mapping](docs/mapping.md): Mapping data in aiflow applications

In addition to low-level APIs used to support and run aiflow applications, the Core library contains some high-level APIs.  There is an API that can be used to programmatically create and run an application.  There are also interfaces that can be implemented to create your own aiflow contributions, such as Triggers and Activities. 

* [Application](docs/app-api.md): API to build and execute a aiflow application
* [Contributions](docs/contribs.md): APIs and interfaces for aiflow contribution development

## Contributing
Want to contribute to Project AIflow? We've made it easy, all you need to do is fork the repository you intend to contribute to, make your changes and create a Pull Request! Once the pull request has been created, you'll be prompted to sign the CLA (Contributor License Agreement) online.

Not sure where to start? No problem, you can browse the Project aiflow repos and look for issues tagged `kind/help-wanted` or `good first issue`. To make this even easier, we've added the links right here too!
* Project AIflow: [kind/help-wanted](https://github.com/TIBCOSoftware/AIflow/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and [good first issue](https://github.com/TIBCOSoftware/AIflow/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
* aiflow cli: [kind/help-wanted](https://github.com/r2d2-ai/ai-box/cli/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and [good first issue](https://github.com/r2d2-ai/ai-box/cli/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
* aiflow core: [kind/help-wanted](https://github.com/r2d2-ai/ai-box/core/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and [good first issue](https://github.com/r2d2-ai/ai-box/core/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
* aiflow contrib: [kind/help-wanted](https://github.com/r2d2-ai/ai-box/contrib/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and [good first issue](https://github.com/r2d2-ai/ai-box/contrib/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)

Another great way to contribute to Project aiflow is to check [AIflow-contrib](https://github.com/r2d2-ai/ai-box/contrib). That repository contains some basic contributions, such as activities, triggers, etc. Perhaps there is something missing? Create a new activity or trigger or fix a bug in an existing activity or trigger.

If you have any questions, feel free to post an issue and tag it as a question, email AIflow-oss@tibco.com or chat with the team and community:

* The [project-AIflow/Lobby](https://gitter.im/r2d2-ai/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link) Gitter channel should be used for general discussions, start here for all things AIflow!
* The [project-AIflow/developers](https://gitter.im/r2d2-ai/developers?utm_source=share-link&utm_medium=link&utm_campaign=share-link) Gitter channel should be used for developer/contributor focused conversations. 

For additional details, refer to the [Contribution Guidelines](https://github.com/TIBCOSoftware/AIflow/blob/master/CONTRIBUTING.md).

## License 
aiflow source code in [this](https://github.com/r2d2-ai/ai-box/core) repository is under a BSD-style license, refer to [LICENSE](https://github.com/r2d2-ai/ai-box/core/blob/master/LICENSE) 
