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
	"strconv"

	"net/http"

	"github.com/googollee/go-socket.io"
)

func loop() {
	fmt.Println("Starting Ansel in AUTONOMOUS mode...")

	sensor.Connect()
	defer sensor.Disconnect()

	// Ultrasonic sensor loop
	go sensor.Loop()

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
}

func server() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		fmt.Println("connected")

		for i := 0; i < 10000; i++ {
			so.Emit("sensor", strconv.Itoa(i), func(so socketio.Socket, data string) {
				fmt.Printf("ack for: " + strconv.Itoa(i))
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

	// Start the webserver
	go server()

	// Start the main control loop
	loop()
}
