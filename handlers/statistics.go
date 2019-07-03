package handlers

import (
	"encoding/json"
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
	var err error
	if http.MethodGet != r.Method {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	statistics, err := h.container.DB.AllStatistics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(statistics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
