#!/bin/bash
d=$(date -R)
nohup ./build/meter -config=build/config.json -group="$1" -user="$2" -last="$3" >>log/test_"$d".log &
