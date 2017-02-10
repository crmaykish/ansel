package main

import "fmt"
import "github.com/crmaykish/ansel/sensor"
import "time"

func main() {
	fmt.Printf("Starting Ansel control program...\n")

	sensor.Connect()
	defer sensor.Disconnect()

	// Ultrasonic sensor loop
	go sensor.Loop()

	for {
		fmt.Println("Main loop forever...")
		time.Sleep(time.Second * 5)
	}
}
