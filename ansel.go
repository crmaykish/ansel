package main

import (
	"fmt"
	"time"

	"os"

	"log"

	"os/signal"
	"syscall"

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

func stop() {
	if !motor.Connected {
		motor.Connect()
	}
	motor.StopMovement()
	motor.Disconnect()
}

func main() {
	// Watch for an OS interupt and trigger a cleanup
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		stop()
		os.Exit(1)
	}()

	// Read command line arguments to determine starting mode
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
