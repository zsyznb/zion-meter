#!/bin/bash

workdir=~/software/hotstuff/zion-meter/pkg/go_abi/stat_abi

mkdir -p $workdir

abigen --sol stat.sol --pkg stat_abi > $workdir/stat_abi.go
