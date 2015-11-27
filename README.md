<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Coporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Plugin - snap publisher to PostgreSQL

[![Build Status](https://api.travis-ci.com/intelsdi-x/snap-plugin-publisher-postgresql.svg?token=FhmCtm9AdqhSXoSbqxo2&branch=master)](https://travis-ci.com/intelsdi-x/snap-plugin-publisher-postgresql)

1. [Getting Started](#getting-started)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3.  [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License and Authors](#license-and-authors)
6. [Thank You](#thank-you)


### Run PostgreSQL server.

```
docker run --name postgres_server -p 5432:5432 --env 'DB_USER=snap' \
--env 'DB_PASS=snap' -d sameersbn/postgresql:9.4-8
```

### Compile plugin
```
make
```

### Export path to the plugin
```
export SNAP_POSTGRESQL_PLUGIN=`pwd`
```
## Documentation

### Examples
Example running psutil plugin, passthru processor, and writing data to a postgresql database.

Documentation for snap collector psutil plugin can be found here (https://github.com/intelsdi-x/snap-plugin-collector-psutil)

In one terminal window, open the snap daemon :
```
$ snapd -l 1
```

In another terminal window:

Load postgresql plugin
```
$ snapctl plugin load $SNAP_POSTGRESQL_PLUGIN/build/rootfs/snap-plugin-publisher-postggresql

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
                "/psutil/load/load1": {},
                "/psutil/load/load15": {}
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

## Example postgresql table


|     time_posted       |     key_column      | value_column  |
|-----------------------|:-------------------:|--------------:|
|2015-09-24 10:06:15+00 | psutil, load, load1 | 1.58          |

### Community Support
This repository is one of **many** plugins in the **snap framework**: a powerful telemetry agent framework.
The full project is at https://github.com/intelsdi-x/snap.

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-postgresql/issues).

## Contributing
We love contributions! :heart_eyes:

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License and Authors
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

## Thank You
And **thank you!** Your contribution is incredibly important to us.
