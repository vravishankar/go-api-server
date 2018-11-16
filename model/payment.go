package model

import "time"

// Payment - Type Definition
type Payment struct {
	TxnID   int       `json:"txn_id"`
	TxnDate time.Time `json:"txn_date"`
}
