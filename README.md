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

# **pulse-plugin-publisher-postgresql**
Pulse Publisher Plugin to PostgreSQL

[![Build Status](https://magnum.travis-ci.com/intelsdi-x/pulse-plugin-publisher-postgresql.svg?token=2ujsxEpZo1issFyVWX29&branch=master)](https://magnum.travis-ci.com/intelsdi-x/pulse-plugin-publisher-postgresql)

## Description
    This plugin publishes data into PostgreSQL for Pulse compliant collectors.

## Dependencies
    It requires project Pulse: https://github.com/intelsdi-x/pulse.

## Configuration
    1. Set PULSE_PATH envoriment variable for running an example.
    2. Change sample configuration in ./examples/psutil-postgresql.json file
    3. Run the example task from the source root directory. E.g. 
    pulsectl task create -t ./examples/psutil-postgresql.json


## Details
    Comma delimitered namespace. E.g.


|     time_posted       |     key_column      | value_column  |
|-----------------------|:-------------------:|--------------:|
|2015-09-24 10:06:15+00 | psutil, load, load1 | 1.58          |

## Change log
    first PR 2015-9-28
