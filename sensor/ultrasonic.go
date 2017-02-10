package sensor

import (
	"log"
	"strconv"
	"strings"

	"bufio"

	"errors"

	"fmt"

	"github.com/tarm/serial"
)

var connected = false
var port *serial.Port
var reader *bufio.Reader

var Data map[int]int

func Connect() {
	fmt.Println("Connecting to Sensor...")
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	var err error
	port, err = serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	} else {
		reader = bufio.NewReader(port)
		Data = make(map[int]int)
		connected = true
	}
}

func Disconnect() {
	fmt.Println("Disconnecting from Sensor...")
	port.Flush()
	port.Close()
	connected = false
}

func Loop(count int) {
	for i := 0; i < count; i++ {
		line, err := readLine()

		if err != nil {
			log.Fatal(err)
		}

		k, v, err := parseSerial(line)

		if err == nil {
			Data[k] = v
		}
		fmt.Printf("%d : %d", k, Data[k])
	}
}

func readLine() (line string, err error) {
	if !connected {
		return "", errors.New("Serial not connected")
	}

	return reader.ReadString('\n')
}

func parseSerial(line string) (key int, value int, err error) {
	p := strings.Split(strings.TrimSpace(line), ":")

	k, err := strconv.Atoi(p[0])
	if err != nil {
		return 0, 0, err
	}

	v, err := strconv.Atoi(p[1])
	if err != nil {
		return 0, 0, err
	}

	return k, v, nil
}
