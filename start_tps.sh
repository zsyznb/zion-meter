#!/bin/bash
d=$(date -R)
exec 1>log/test_"$d".log
build/./meter -config=build/config.json -group="$1" -user="$2" -last="$3"
