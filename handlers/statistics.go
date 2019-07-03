package handlers

import (
	"github.com/benja-M-1/thegoodcorner/app"
	"net/http"
)

type StatisticsHandler struct {
	container *app.Container
}

func NewStatisticsHandler(c *app.Container) StatisticsHandler {
	h := StatisticsHandler{c}

	return h
}

func (h *StatisticsHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
