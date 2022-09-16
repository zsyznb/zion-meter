#!/bin/bash
chmod  777 ./test.log
date=$(date -R)
 mv ./test.log ./log/tpslog_group="$1"_user="$2"_last="$3"_time="$date"
 # shellcheck disable=SC2028
 echo "Save log successfully! Time : $date"
