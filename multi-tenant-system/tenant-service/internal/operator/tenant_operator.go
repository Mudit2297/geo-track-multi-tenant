package operator

import (
	"database/sql"
	"fmt"
	"tenant-service/internal/model"
)

type TenantOperator struct {
	DB *sql.DB
}

func NewTenantOperator(db *sql.DB) *TenantOperator {
	return &TenantOperator{DB: db}
}

func (r *TenantOperator) CreateTenant(tenant model.Tenant) error {
	_, err := r.DB.Exec(`INSERT INTO tenants (name, contact_email) VALUES ($1, $2)`, tenant.Name, tenant.ContactEmail)
	return err
}

func (r *TenantOperator) GetTenantByID(id string) (model.Tenant, error) {
	var t model.Tenant
	err := r.DB.QueryRow(`SELECT id, name, contact_email from tenants WHERE id = $1`, id).Scan(&t.ID, &t.Name, &t.ContactEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Tenant{}, fmt.Errorf("no rows returned - %w", err)
		}
		return model.Tenant{}, err
	}

	return t, nil
}

func (r *TenantOperator) GetAllTenants() ([]model.Tenant, error) {
	rows, err := r.DB.Query(`SELECT id, name, contact_email FROM tenants`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []model.Tenant
	for rows.Next() {
		var t model.Tenant
		err := rows.Scan(&t.ID, &t.Name, &t.ContactEmail)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}
