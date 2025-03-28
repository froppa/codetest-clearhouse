package repositories

import (
	"github.com/froppa/company-api/internal/models"
	"github.com/google/uuid"
)

type Company struct {
	ID      *uuid.UUID `db:"id,omitempty"`
	Name    string     `db:"name"`
	Country string     `db:"country"`
	Email   string     `db:"email"`
}

type Owner struct {
	ID        *uuid.UUID `db:"id,omitempty"`
	CompanyID *uuid.UUID `db:"company_id,omitempty"`
	Name      string     `db:"name"`
	SSN       string     `db:"ssn"`
}

func (c Company) toInternal() models.Company {
	return models.Company{
		ID:      c.ID,
		Name:    c.Name,
		Country: c.Country,
		Email:   c.Email,
	}
}

func (o Owner) toInternal() models.Owner {
	return models.Owner{
		ID:        o.ID,
		CompanyID: o.CompanyID,
		Name:      o.Name,
		SSN:       o.SSN,
	}
}

func fromInternal(in models.Company) Company {
	return Company{
		ID:      in.ID,
		Name:    in.Name,
		Country: in.Country,
		Email:   in.Email,
	}
}

func fromInternalOwner(in models.Owner) Owner {
	return Owner{
		ID:        in.ID,
		Name:      in.Name,
		CompanyID: in.CompanyID,
		SSN:       in.SSN,
	}
}
