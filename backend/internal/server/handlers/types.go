package handlers

import "net/http"

type Handler interface {
	Controller(w http.ResponseWriter, r *http.Request)
}

type Message struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type WebsocketHandler struct {
	Queue chan<- Message
}
