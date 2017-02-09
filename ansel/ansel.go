package main

import (
	"fmt"
	"strconv"

	"github.com/tarm/serial"

	"bufio"
	"log"
	"strings"
)

var sensorData map[int]int

func main() {
	fmt.Printf("Starting Ansel control program...\n")

	sensorData = make(map[int]int)

	// TODO: add read timeout
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)

	for {
		reading, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		p := strings.Split(strings.TrimSpace(reading), ":")

		k, err := strconv.Atoi(p[0])
		if err != nil {
			fmt.Printf("Error parsing key: %v\n", k)
		}

		v, err := strconv.Atoi(p[1])
		if err != nil {
			fmt.Printf("Error parsing value: %v\n", v)
		}

		sensorData[k] = v

		fmt.Printf("%d : %d\n", k, sensorData[k])

	}

}
