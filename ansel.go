package main

import (
	"fmt"
	"time"

	"os"

	"os/signal"
	"syscall"

	"github.com/crmaykish/ansel/motor"
	"github.com/crmaykish/ansel/sensor"
	"github.com/crmaykish/ansel/server"
)

// Really basic obstacle avoid algorithm
func control() {
	for {
		f := sensor.Data(sensor.Front)
		l := sensor.Data(sensor.FrontLeft)
		r := sensor.Data(sensor.FrontRight)

		if f <= sensor.SafeDistance && sensor.SafeDistance >= 0 || l <= sensor.SafeDistance && sensor.SafeDistance >= 0 || r <= sensor.SafeDistance && sensor.SafeDistance >= 0 {
			var dir string
			if l > r || l < 0 {
				dir = "left"
			} else {
				dir = "right"
			}

			fmt.Println("Turning " + dir)

			motor.SetMovement(dir, 190)
		} else {
			fmt.Println("Forward")
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

	if sensor.Connected {
		sensor.Disconnect()
	}
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

	// Start the sensor thread
	sensor.Connect()
	go sensor.Loop()

	// Connect to motors
	motor.Connect()

	// Start control loop thread
	go control()

	// Start the webserver on the main thread
	server.Start()
}
