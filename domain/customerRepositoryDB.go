package domain

import (
	"banking/logger"
	"banking/utils"
	"context"

	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (db *CustomerRepositoryDB) Create(customer *Customer) (int, error) {
	result, err := db.client.Exec(`INSERT INTO customer
		(name, city, zipcode)
		VALUES(?, ?, ?)`,
		customer.Fullname,
		customer.City,
		customer.Zipcode)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	customerId, err := result.LastInsertId()
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return int(customerId), nil
}

func (db *CustomerRepositoryDB) Get(customerID int) (*Customer, error) {
	var customer *Customer
	err := db.client.Select(customer,
		`SELECT id, name, city, zipcode
		FROM customer
		WHERE id = ?`,
		customerID)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	return customer, nil
}

// contructor
func NewCustomerRepositoryDB(ctx *context.Context) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}