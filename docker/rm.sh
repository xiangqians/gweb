#!/bin/bash

# 停止容器
docker stop gweb

# 删除容器
docker rm gweb

# 删除镜像
docker rmi gweb:1.0.0
