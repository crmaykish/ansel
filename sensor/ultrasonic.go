package sensor

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bufio"

	"errors"

	"github.com/tarm/serial"
)

var connected = false
var port *serial.Port
var reader *bufio.Reader

var Data map[int]int

func Connect() {
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	var err error
	port, err = serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	} else {
		reader = bufio.NewReader(port)
		connected = true
	}
}

func Disconnect() {
	port.Flush()
	port.Close()
	connected = false
}

func readLine() (line string, err error) {
	if !connected {
		return "", errors.New("Serial not connected")
	}

	return reader.ReadString('\n')
}

func parseSerial(line string) (key, value int) {
	p := strings.Split(strings.TrimSpace(line), ":")

	k, err := strconv.Atoi(p[0])
	if err != nil {
		fmt.Printf("Error parsing key: %v\n", k)
	}

	v, err := strconv.Atoi(p[1])
	if err != nil {
		fmt.Printf("Error parsing value: %v\n", v)
	}

	return k, v
}
