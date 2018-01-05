#!/bin/bash 

echo "create preset dirs..."

project="gocrawler"

rm -rf $project
mkdir -p $project/bin

echo "copy resource..."
cp -r assets $project > /dev/null

echo "build..."
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o $project/bin/$project

echo "upload to server..."
scp -r $project root@10.0.2.206:/data/go

#修改权限
ssh -t -p 22 root@10.0.2.206 "chmod -R a+x /data/go/$project"

ssh -t -p 22 root@10.0.2.206 "mkdir /data/go/$project/data" > /dev/null

#删除本地缓存文件
rm -rf $project

echo "finished!"

#窗口不自动消失
`read`