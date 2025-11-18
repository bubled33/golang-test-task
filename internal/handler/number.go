package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo Repository
}

type Repository interface {
	AddNumber(ctx context.Context, num int) error
	GetAllSorted(ctx context.Context) ([]int, error)
}

type Request struct {
	Number int `json:"number"`
}

func New(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) AddNumber(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.AddNumber(r.Context(), req.Number); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	numbers, err := h.repo.GetAllSorted(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}
