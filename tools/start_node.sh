#!/bin/bash

file=$GOPATH/src/github.com/Vedaad-Shakib/IoTeX-Sim/sampleconfig/stonevan/config_$1.yaml
echo "Starting node with config file:" $file
$GOPATH/src/github.com/Vedaad-Shakib/IoTeX-Sim/bin/server -stderrthreshold=WARNING -log_dir=./log -config=$file

exit 0