#!/bin/bash
if [ "$1" == "full" ]; then
    echo "Cross-compiling Ansel for Linux on ARM"
    env GOOS=linux GOARCH=arm GOARM=7 go build -o /c/go_ws/bin/arm/ansel ansel.go

    echo "Copying binary to Ansel..."
    scp -q /c/go_ws/bin/arm/ansel colin@ansel:/home/colin
fi

echo "Copying web assets to Ansel..."
scp -rq ./assets/ colin@ansel:/home/colin/

echo "Done!"