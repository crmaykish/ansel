package main

import (
	"fmt"
	"time"

	"os"

	"log"

	"github.com/crmaykish/ansel/motor"
	"github.com/crmaykish/ansel/sensor"
)

func interactive() {
	fmt.Println("Starting Ansel in INTERACTIVE mode...")
	// TODO: Interactive driving loop
}

func autonomous() {
	fmt.Println("Starting Ansel in AUTONOMOUS mode...")

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

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "i" {
			interactive()
		} else {
			log.Fatal("Unrecognized argument: " + args[0])
		}
	} else {
		autonomous()
	}
}
