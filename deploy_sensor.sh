#!/bin/bash
serial_port=/dev/ttyUSB1;
board_model=nano328;


cd ultrasonic/;
echo "Building Ansel Ultrasonic Sensor Program for board:" $board_model
sudo ino build -m $board_model;
echo "Uploading to" $board_model "on" $serial_port
sudo ino upload -m $board_model -p $serial_port;
cd ../;
