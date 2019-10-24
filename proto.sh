#!/bin/bash
#遍历src目录，执行所有的proto文件
function read_dir() {
  # shellcheck disable=SC2045
  for file in $(ls $1); do
    if [ -d $1"/"$file ]; then
      read_dir $1"/"$file
    else
      if [ "${file##*.}"x = "proto"x ]; then
        echo $1"/"$file
        protoc --go_out=plugins=grpc:. $1"/"$file
      fi
    fi
  done
}
read_dir ./src