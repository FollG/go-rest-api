package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetList)
}

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user list"))
}
