package models

import "time"

type DrugDoc struct {
	TradeName   string     `json:"trade_name" form:"trade_name"`
	DrugName    string     `json:"drug_name" form:"drug_name"`
	Description string     `json:"description" form:"description"`
	Preparation string     `json:"preparation" form:"preparation"`
	Caution     string     `json:"caution" form:"caution"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type DrugDocWithID struct {
	ID string `json:"id" form:"id"`
	DrugDoc
}
