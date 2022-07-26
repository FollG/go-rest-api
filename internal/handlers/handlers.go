package handler

import "github.com/julienschmidt/httprouter"

type Handler interface {
	Register(route *httprouter.Router)
}
