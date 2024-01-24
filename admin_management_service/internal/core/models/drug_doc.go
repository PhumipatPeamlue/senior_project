package models

import "time"

type DrugDoc struct {
	ID          string     `json:"id"`
	TradeName   string     `json:"trade_name"`
	DrugName    string     `json:"drug_name"`
	Description string     `json:"description"`
	Preparation string     `json:"preparation"`
	Caution     string     `json:"caution"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
