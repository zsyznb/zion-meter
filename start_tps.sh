#!/bin/bash
d=$(date -R)
exec 1>log/test_"$d".log
build/local/./meter -config=build/local/config.json -group="$1" -user="$2" -last="$3"
