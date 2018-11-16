package repository

import (
	"fmt"
	"time"

	"../model"
)

// GetAll - Get All Payments
func GetAll() ([]model.Payment, error) {
	var payments []model.Payment
	m := map[string]interface{}{}

	fmt.Println("Fetching payments from the database.")
	query := "SELECT * FROM paytxns"
	iterable := Session.Query(query).Iter()
	for iterable.MapScan(m) {
		payments = append(payments, model.Payment{
			TxnID:   m["txn_id"].(int),
			TxnDate: m["txn_date"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	return payments, nil
}
