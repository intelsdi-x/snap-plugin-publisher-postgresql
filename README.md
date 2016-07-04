# snap publisher plugin - PostgreSQL

[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-publisher-postgresql.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-postgresql)
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-publisher-postgresql)](http://goreportcard.com/report/intelsdi-x/snap-plugin-publisher-postgresql)



1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
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

#### Compile plugin
```
make
```

#### Export path to the plugin
```
export SNAP_POSTGRESQL_PLUGIN=`pwd`
```

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation

<< @TODO

### Examples

Example running psutil plugin, passthru processor, and writing data to a postgresql database.

Documentation for snap collector psutil plugin can be found here (https://github.com/intelsdi-x/snap-plugin-collector-psutil)

In one terminal window, open the snap daemon :
```
$ snapd -l 1 -t 0
```

In another terminal window:

Load postgresql plugin
```
$ snapctl plugin load $SNAP_POSTGRESQL_PLUGIN/build/rootfs/snap-plugin-publisher-postgresql

```

Load psutil plugin
```
$ snapctl plugin load $SNAP_PSUTIL_PLUGIN/build/rootfs/snap-plugin-collector-psutil
```

See available metrics for your system
```
$ snapctl metric list
```

Create a task JSON file for example sample-task.json:    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/psutil/load/load1": {},
                "/intel/psutil/load/load15": {}
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
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
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Load passthru plugin for processing:
```
$ snapctl plugin load build/rootfs/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:44:03 PST
```

Create task:
```
$ snapctl task create -t sample-task.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
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
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions! 

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
