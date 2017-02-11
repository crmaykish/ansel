package main

import (
	"fmt"
	"time"

	"os"

	"os/signal"
	"syscall"

	"strconv"

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
		f := sensor.Data[sensor.Front]
		l := sensor.Data[sensor.FrontLeft]
		r := sensor.Data[sensor.FrontRight]

		// TODO: add checks for left and right too

		if f <= sensor.SafeDistance && sensor.SafeDistance >= 0 {
			fmt.Println("Front is TOO CLOSE, L: " + strconv.Itoa(l) + ", R: " + strconv.Itoa(r))

			var dir string
			if l > r || l < 0 {
				dir = "left"
			} else {
				dir = "right"
			}

			motor.SetMovement(dir, 255)
		} else {
			fmt.Println("Front is a safe distance")
			motor.SetMovement("forward", 255)
		}

		// TODO: this should be a timer with an elapsed time check, not just a fixed delay
		time.Sleep(motor.UpdateDelay)
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
		switch args[0] {
		case "i":
			interactive()
		case "s":
			stop()
		case "?":
			fmt.Printf("i : Interactive mode\n" + "s : Stop motors\n" + "? : This help screen\n")
		default:
			fmt.Printf("Unrecognized argument: " + args[0] + "\nRun ./ansel ? to see options\n")
		}
	} else {
		autonomous()
	}
}
