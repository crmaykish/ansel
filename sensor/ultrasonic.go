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

const Front = 5
const FrontRight = 6
const RightFront = 7
const RightRear = 8
const RearRight = 9
const Rear = 0
const RearLeft = 1
const LeftRear = 2
const LeftFront = 3
const FrontLeft = 4

const SafeDistance = 30

var Data map[int]int

var connected = false
var port *serial.Port
var reader *bufio.Reader

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

func Loop() {
	for {
		line, err := readLine()

		if err != nil {
			log.Fatal(err)
		}

		k, v, err := parseSerial(line)

		if err == nil {
			Data[k] = v
		}
	}
}

func readLine() (line string, err error) {
	if !connected {
		return "", errors.New("Serial not connected")
	}

	return reader.ReadString('\n')
}

func parseSerial(line string) (key int, value int, err error) {
	if !strings.Contains(line, ":") {
		return 0, 0, errors.New("Bad serial data")
	}
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

func PrintData() {
	var d string
	for i := 0; i < 10; i++ {
		d += strconv.Itoa(Data[i])
		if i != 9 {
			d += " | "
		}
	}
	fmt.Println(d)
}
