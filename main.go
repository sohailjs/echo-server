package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("connection successful")

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		strMsg := bytes.NewBuffer(msg).String()
		fmt.Println(strMsg)

		serverMsg := []byte("this is server: " + strMsg)
		conn.WriteMessage(mt, serverMsg)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/chat", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
