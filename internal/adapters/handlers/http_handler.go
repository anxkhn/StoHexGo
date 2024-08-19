package handlers

import (
	"encoding/json"
	"net/http"
	"stock-portfolio/internal/core/domain"
	"stock-portfolio/internal/core/ports"
)

type HTTPHandler struct {
	portfolioService ports.PortfolioService
}

func NewHTTPHandler(portfolioService ports.PortfolioService) *HTTPHandler {
	return &HTTPHandler{portfolioService: portfolioService}
}

func (h *HTTPHandler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"msg": "🚀 Server is running!"}
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) HandleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Transactions []domain.Transaction `json:"transactions"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profit, err := h.portfolioService.ExecuteTransactions(req.Transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"profit": profit})
}

func (h *HTTPHandler) HandleBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	balance, err := h.portfolioService.GetCurrentBalance()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(balance)
}
