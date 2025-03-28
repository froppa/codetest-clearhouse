package services

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

// SSNValidationRequest represents the request body for SSN validation
type SSNValidationRequest struct {
	SSN string `json:"ssn"`
}

// ValidateSSN validates the SSN format
func (s SSNValidationRequest) validFormat() bool {
	re := regexp.MustCompile(`^\d{6}-\d{4}$`)
	return re.MatchString(s.SSN)
}

var accessGroups = map[string][]string{
	"admin":      {"read", "write", "validateSSN"},
	"restricted": {"validateSSN"},
}

// SSNValidationHandler validates the SSN from the request body
func SSNValidationHandler(w http.ResponseWriter, r *http.Request) {
	var request SSNValidationRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !request.validFormat() {
		http.Error(w, "Invalid SSN Format", http.StatusBadRequest)
		return
	}

	// Simulate SSN validation
	isValid := func() bool {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return r.Intn(2) == 0 // 50%-50%
	}

	// Only return the SSN status for authorized access groups
	group := r.URL.Query().Get("group")
	if group == "restricted" {
		// If restricted access, return a generic response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Access granted. SSN validation completed."})
		return
	}

	// For groups with permission to read SSN, return the result
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ssn":     request.SSN,
		"valid":   isValid(),
		"message": "SSN validation complete.",
	})
	w.WriteHeader(http.StatusOK)
}

// Middleware to check for required permission
func MiddlewareCheckPermission(requiredPermission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// In a real app, you'd check the user from the context or headers
			// Here we'll simulate with a query parameter `group`
			group := r.URL.Query().Get("group")
			if group == "" {
				http.Error(w, "Access group not provided", http.StatusForbidden)
				return
			}

			permissions, exists := accessGroups[group]
			if !exists || !contains(permissions, requiredPermission) {
				http.Error(w, "You do not have the required permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func contains(permissions []string, permission string) bool {
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}
