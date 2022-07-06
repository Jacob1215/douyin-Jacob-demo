#! /bin/bash

if [ ! -f 'main' ]; then
  echo "文件不存在! 待添加的安装包:" 'main'
  exit
fi

echo "demo-go..."
sleep 3
docker stop demo-go

sleep 2
docker rm demo-go

docker rmi demo-go
echo ""

echo "deomo-go packaging..."
sleep 3
docker build -t demo-go .
echo ""

echo "demo-go running..."
sleep 3
docker run --name demo-go \
  -p 9801:9801 \
  -d demo-go

docker logs -f demo-go | sed '/Started CashierApplication/q'

echo ""