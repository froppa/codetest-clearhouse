package repositories

import (
	"context"

	"github.com/froppa/company-api/internal/models"
	"github.com/upper/db/v4"
)

func Companies(sess db.Session) db.Collection {
	return sess.Collection("companies")
}

func (r *Repository) CreateCompany(ctx context.Context, company models.Company) (models.Company, error) {
	model := fromInternal(company)
	err := Companies(r.db).InsertReturning(&model)
	if err != nil {
		return models.Company{}, err
	}

	return model.toInternal(), nil
}

func (r *Repository) ListCompanies(ctx context.Context) ([]models.Company, error) {
	var model []Company
	if err := Companies(r.db).Find().All(&model); err != nil {
		return nil, err
	}

	companies := make([]models.Company, len(model))
	for _, company := range model {
		companies = append(companies, company.toInternal())
	}

	return companies, nil
}

func (r *Repository) GetCompanyByID(ctx context.Context, companyID string) (models.Company, error) {
	var company Company
	if err := r.db.Collection("companies").Find("id", companyID).One(&company); err != nil {
		return models.Company{}, err
	}

	return company.toInternal(), nil
}

func (r *Repository) AddCompanyOwner(ctx context.Context, companyID string, owner models.Owner) (models.Owner, error) {
	model := fromInternalOwner(owner)

	err := r.db.Collection("owners").InsertReturning(&model)
	if err != nil {
		return models.Owner{}, nil
	}

	return model.toInternal(), nil
}
