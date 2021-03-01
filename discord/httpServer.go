package main

import (
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Data    string `json:"data"`
	Channel string `json:"channel"`
	Login   string `json:"login"`
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	switch r.Method {
	case "POST":
		_, err := r.Body.Read(data)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		_, _ = gBot.ChannelMessageSend("778637653406122004", string(data))
		w.WriteHeader(200)
	}
}

func startServer() {
	http.HandleFunc("/discord", sendHandler)
	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
