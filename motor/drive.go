package motor

import (
	"fmt"
	"log"
	"strconv"

	"time"

	"github.com/tarm/serial"
)

// UpdateRate specifies the rate (in Hertz) to send motor board messages
const UpdateRate = 20

// UpdateDelay is the time between motor loop iterations
const UpdateDelay = time.Second / UpdateRate

var port *serial.Port
var connected = false

// Connect to the serial port
func Connect() {
	fmt.Println("Connecting to Motors...")
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	var err error
	port, err = serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	} else {
		connected = true
	}
}

// Disconnect closes the serial port
func Disconnect() {
	fmt.Println("Disconnecting from Sensor...")
	port.Flush()
	port.Close()
	connected = false
}

func sendSerial(message string) {
	port.Write([]byte(message + "\n"))
}

func SetMovement(direction string, speed int) {
	sendSerial("d," + direction + "," + strconv.Itoa(speed))
}

func StopMovement() {
	sendSerial("stop,")
}
