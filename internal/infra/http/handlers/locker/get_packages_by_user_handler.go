package locker

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	lockerusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/locker"
)

// GetPackagesByUserHandler handles requests to get packages by user
type GetPackagesByUserHandler struct {
	getPackagesByUserCase *lockerusecases.GetPackagesByUserCase
}

// NewGetPackagesByUserHandler creates a new instance of GetPackagesByUserHandler
func NewGetPackagesByUserHandler(getPackagesByUserCase *lockerusecases.GetPackagesByUserCase) *GetPackagesByUserHandler {
	return &GetPackagesByUserHandler{
		getPackagesByUserCase: getPackagesByUserCase,
	}
}

// Handle handles the HTTP request to get packages by user
func (h *GetPackagesByUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Get user ID from query parameter
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user_id format", http.StatusBadRequest)
		return
	}

	// Get packages
	packages, err := h.getPackagesByUserCase.Execute(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return packages as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages)
}
