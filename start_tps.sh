#!/bin/bash
exec 1>test.log
build/local/./meter -config=build/local/config.json -group="$1" -user="$2" -last="$3"
