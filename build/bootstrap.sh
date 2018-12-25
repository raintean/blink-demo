#!/usr/bin/env bash

cd /app

echo -e "----------\ncleaning...\n----------"
rm -rf ./dist
rm -rf ./ui/dist

echo -e "----------\nbuild web ui...\n----------"
cd /app/ui
cnpm install
npm run build

echo -e "----------\ngenerate bin data file...\n----------"
go-bindata -o bin/ui_bin.go -pkg ui -prefix ./dist ./dist/...

echo -e "----------\nbuild go application...\n----------"
if [ "${GOARCH}" = "386" ];then
    export CC=i686-w64-mingw32-gcc
else
    export CC=x86_64-w64-mingw32-gcc
fi

export GOPATH=/app/.go

cd /app
go build \
    -tags="bdebug" \
    -ldflags="-H=windowsgui" \
    -o dist/blink-demo.exe \
    blink-demo/cmd/main

echo -e "----------\ncompress executable file...\n----------"
upx dist/blink-demo.exe


