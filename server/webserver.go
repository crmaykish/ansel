package server

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/crmaykish/ansel/sensor"
	"github.com/googollee/go-socket.io"
)

var count = 0

var Server *socketio.Server

func Start() {
	var err error
	Server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	Server.On("connection", func(so socketio.Socket) {
		fmt.Println("Client connected")
		count++

		for {
			fmt.Println("broadcasting for ", count)
			fmt.Println(sensor.Json())
			so.Emit("sensor", sensor.Json(), nil)
			time.Sleep(time.Millisecond * 200)
		}

	})

	Server.On("error", func(so socketio.Socket, err error) {
		log.Println("Error: ", err)
	})

	http.Handle("/socket.io/", Server)
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	fmt.Println("Starting web server")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
