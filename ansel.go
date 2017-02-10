package main

import "fmt"
import "github.com/crmaykish/ansel/sensor"

func main() {
	fmt.Printf("Starting Ansel control program...\n")

	defer sensor.Disconnect()

	sensor.Connect()
}
