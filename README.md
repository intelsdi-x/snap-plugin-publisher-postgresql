DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


# Snap publisher plugin - PostgreSQL

[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-publisher-postgresql.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-postgresql)
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-publisher-postgresql)](http://goreportcard.com/report/intelsdi-x/snap-plugin-publisher-postgresql)



1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Task Manifest Config](#task-manifest-config)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

### Installation

#### Run PostgreSQL server.

```
docker run --name postgres_server -p 5432:5432 --env 'DB_USER=snap' \
--env 'DB_PASS=snap' -d sameersbn/postgresql:9.4-8
```

#### Download File plugin binary:
You can get the pre-built binaries for your OS and architecture at plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-publisher-postgresql/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-publisher-postgresql

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-postgresql.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
### Task Manifest Config

In task manifest, the config section of PostgreSQL publisher describes how to establish a connection to the PostgreSQL server.

Name | Data Type | Description
----------|-----------|---------------|-------------
hostname | string | the host of PostgreSQL service
port | number | the port number of PostgreSQL service
username | string | the name of user
password | string | the password of user
database | string | the name of database 
table_name | string | the name of table

### Examples

Example of running [psutil collector plugin](https://github.com/intelsdi-x/snap-plugin-collector-psutil) and publishing data to PostgreSQL database.

Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `sudo snapteld -l 1 -t 0 &`


Download and load Snap plugins (paths to binary files for Linux/amd64):
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-postgresql/latest/linux/x86_64/snap-plugin-publisher-postgresql
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-psutil/latest/linux/x86_64/snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-postgresql
$ snaptel plugin load snap-plugin-collector-psutil
```

Create a [task manifest](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md) (see [exemplary tasks](examples/)),
for example `psutil-postgresql.json` with following content:
```json
{
  "version": 1,
  "schedule": {
    "type": "simple",
    "interval": "10s"
  },
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/psutil/load/load1": {},
        "/intel/psutil/load/load15": {}
      },
      "publish": [
        {
          "plugin_name": "postgresql",
          "config": {
            "hostname": "localhost",
            "port": 5432,
            "username": "snap",
            "table_name": "snap",
            "database": "snap",
            "password": "snap"
          }
        }
      ]
    }
  }
}
```

Create a task:
```
$ snaptel task create -t psutil-postgresql.json
```

Watch created task:
```
$ snaptel task watch <task_id>
```

To stop previously created task:
```
$ snaptel task stop <task_id>
```

Example postgresql table

|     time_posted       |     key_column           | value_column  |
|-----------------------|:------------------------:|--------------:|
|2015-09-24 10:06:15+00 | intel.psutil.load.load1  | 1.58          |
|2015-09-24 10:06:15+00 | intel.psutil.load.load15 | 2.43          |

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-postgresql/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-postgresql/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-publisher-postgresql/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions! 

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
