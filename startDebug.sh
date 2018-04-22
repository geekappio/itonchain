#!/bin/bash

nohup go build -o itonchain -gcflags='-N -l' github.com/geekappio/itonchain/app && dlv --listen=:9000 --headless=true --api-version=2 exec ./itonchain &