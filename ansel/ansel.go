package main

import (
	"fmt"

	"github.com/tarm/serial"

	"bufio"
	"log"
)

func main() {
	fmt.Printf("Starting Ansel control program...\n")

	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)

	for {
		reply, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(reply)
	}

}
