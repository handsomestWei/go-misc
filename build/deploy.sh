#!/bin/bash

image_name=xxx

export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY=https://goproxy.io

## df -h
## mkdir xxx/mygo
## mount -t cifs -o username=Everyone,uid=root,gid=root //172.16.xx.xx/gopath /xxx/mygo
## umount -l /xxx/mygo

go build -o app

if [ $? -ne 0 ]; then
    echo "build failed"
    exit 1
fi

docker login xxx -u xx -p 123456
docker build -t image_name .
docker push image_name
