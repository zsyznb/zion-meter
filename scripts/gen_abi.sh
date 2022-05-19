#!/bin/bash

basedir=~/software/hotstuff/zion-meter/pkg/go_abi
list=(stat data_stat)

for fn in ${list[*]}; do
    abiname=${fn}_abi
    workdir=$basedir/$abiname
    mkdir -p $workdir
    abigen --sol ${fn}.sol --pkg $abiname > $workdir/${abiname}.go
done
