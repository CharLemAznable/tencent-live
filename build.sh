#!/usr/bin/env bash

go build -ldflags="-s -w" -o tencent-live.bin
upx --brute tencent-live.bin
