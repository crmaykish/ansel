package main

import (
	"fmt"
	"time"

	"os"

	"os/signal"
	"syscall"

	"github.com/crmaykish/ansel/motor"
	"github.com/crmaykish/ansel/sensor"

	"log"

	"net/http"

	"github.com/googollee/go-socket.io"
)

func motors() {
	motor.Connect()

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

			motor.SetMovement(dir, 255)
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

// TODO: abstract the server out and let other systems emit events instead of the server "polling" for them
func server() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		fmt.Println("connected")

		for {
			so.Emit("sensor", sensor.Json(), func(so socketio.Socket, data string) {
			})
			time.Sleep(time.Millisecond * 200)
		}

	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	fmt.Println("Starting web server")
	log.Fatal(http.ListenAndServe(":8000", nil))
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

	// Start the webserver thread
	// go server()

	// Start the sensor thread
	sensor.Connect()
	go sensor.Loop()

	// Start the motor control loop thread
	// motors()
	for {
	}
}
