package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/froppa/company-api/internal/models"
	"github.com/froppa/company-api/internal/repositories"
	"github.com/froppa/company-api/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	repo   *repositories.Repository
	logger *zap.Logger
}

func NewHandler(repo *repositories.Repository, logger *zap.Logger) *Handler {
	return &Handler{repo: repo, logger: logger}
}

// RegisterRoutes attaches handlers to router
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/companies", h.ListCompanies).Methods("GET")
	router.HandleFunc("/companies", h.CreateCompany).Methods("POST")
	router.HandleFunc("/companies/{id}", h.GetCompanyByID).Methods("GET")
	router.HandleFunc("/companies/{id}/owners", h.AddOwner).Methods("POST")

	router.Handle(
		"/auth/verify",
		services.
			MiddlewareCheckPermission("validateSSN")(http.HandlerFunc(services.SSNValidationHandler)),
	).Methods("POST")
}

// CreateCompany creates a company (POST /companies)
func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	createdCompany, err := h.repo.CreateCompany(r.Context(), company)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating company: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // This needs to be called before writing the body

	json.NewEncoder(w).Encode(createdCompany)
}

// GetCompanyByID retrieves a single company by ID (GET /companies/:id)
func (h *Handler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	company, err := h.repo.GetCompanyByID(r.Context(), companyID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching company: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(company)
}

// AddOwner adds an owner to an existing company (POST /companies/:id/owners)
func (h *Handler) AddOwner(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	var owner models.Owner
	if err := json.NewDecoder(r.Body).Decode(&owner); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	createdOwner, err := h.repo.AddCompanyOwner(r.Context(), companyID, owner)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding owner: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdOwner)
}

// ListCompanies fetches all companies (GET /companies)
func (h *Handler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.repo.ListCompanies(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching companies: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(companies)
}
