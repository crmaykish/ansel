package main

import (
	"fmt"
	"time"

	"github.com/crmaykish/ansel/motor"
	"github.com/crmaykish/ansel/sensor"
)

func main() {
	fmt.Printf("Starting Ansel control program...\n")

	sensor.Connect()
	defer sensor.Disconnect()

	// Ultrasonic sensor loop
	go sensor.Loop()

	motor.Connect()

	for {
		fmt.Println("Motor loop")
		motor.SetMovement("left", 255)
		time.Sleep(time.Second * 1)
		motor.SetMovement("right", 255)
		time.Sleep(time.Second * 1)
	}
}
