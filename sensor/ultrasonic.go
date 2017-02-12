package sensor

import (
	"log"
	"strconv"
	"strings"

	"bufio"

	"errors"

	"fmt"

	"sync"

	"encoding/json"

	"github.com/tarm/serial"
)

// TODO: wrap up sensors in a struct with easy naming and eror states

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

var dataMutex sync.Mutex
var data map[int]int

var Connected = false
var port *serial.Port
var reader *bufio.Reader

func Data(key int) (val int) {
	dataMutex.Lock()
	v := data[key]
	dataMutex.Unlock()
	return v
}

func SetData(key, val int) {
	dataMutex.Lock()
	data[key] = val
	dataMutex.Unlock()
}

func Connect() {
	fmt.Println("Connecting to Sensor...")
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	var err error
	port, err = serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	} else {
		reader = bufio.NewReader(port)
		data = make(map[int]int)
		Connected = true
	}
}

func Disconnect() {
	fmt.Println("Disconnecting from Sensor...")
	port.Flush()
	port.Close()
	Connected = false
}

func Loop() {
	for {
		line, err := readLine()

		if err != nil {
			log.Fatal(err)
		}

		k, v, err := parseSerial(line)

		if err == nil {
			SetData(k, v)
		}
	}
}

func readLine() (line string, err error) {
	if !Connected {
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

func Json() string {
	dataMutex.Lock()

	s := make(map[string]string)

	for i := 0; i < 10; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(data[i])
		s[k] = v
	}

	dataMutex.Unlock()

	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}

	return string(b)
}
