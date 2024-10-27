package handler

import "net/http"

type Handler struct{}

func (h Handler) HandlerGetProduct(w http.ResponseWriter, r *http.Request) {}

func NewHandler() Handler {
	return Handler{}
}
