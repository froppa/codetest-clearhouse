package models

import "github.com/google/uuid"

// Company represents a company in JSON
type Company struct {
	ID      *uuid.UUID `json:"id,omitempty"`
	Name    string     `json:"name"`
	Country string     `json:"country"`
	Email   string     `json:"email,omitempty"`
}

// Owner represents an owner of a company
type Owner struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	CompanyID *uuid.UUID `json:"company_id"`
	Name      string     `json:"name"`
	SSN       string     `json:"ssn"`
}
