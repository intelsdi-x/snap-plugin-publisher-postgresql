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
