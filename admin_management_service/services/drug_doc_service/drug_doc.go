package drug_doc_service

import "time"

type DrugDoc struct {
	ID          string     `json:"id"`
	TradeName   string     `json:"trade_name"`
	DrugName    string     `json:"drug_name"`
	Description string     `json:"description"`
	Preparation string     `json:"preparation"`
	Caution     string     `json:"caution"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type GetResponse struct {
	Doc      DrugDoc `json:"doc"`
	ImageURL string  `json:"image_url"`
}

type SearchResponse struct {
	Total int       `json:"total"`
	Data  []DrugDoc `json:"data"`
}

type NewDrugDocRequest struct {
	TradeName   string `form:"trade_name"`
	DrugName    string `form:"drug_name"`
	Description string `form:"description"`
	Preparation string `form:"preparation"`
	Caution     string `form:"caution"`
}

type UpdateDrugDocRequest struct {
	ID          string `form:"id"`
	TradeName   string `form:"trade_name"`
	DrugName    string `form:"drug_name"`
	Description string `form:"description"`
	Preparation string `form:"preparation"`
	Caution     string `form:"caution"`
	DeleteImage bool   `form:"delete_image"`
}
