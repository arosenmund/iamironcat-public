#! /bin/bash

echo "Generating new payload for $1 \n"


echo "Payload is set for Windows x64"

export GOOS=windows
export GOARCH=amd64

garble -tiny -literals build -ldflags="-s -w" $1

echo "Created new file\n"

echo "Hosting files with http server on port 8000"

python3 -m http.server

