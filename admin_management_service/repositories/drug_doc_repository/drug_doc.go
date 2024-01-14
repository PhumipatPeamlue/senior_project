package drug_doc_repository

import "time"

type DrugDoc struct {
	TradeName   string     `json:"trade_name"`
	DrugName    string     `json:"drug_name"`
	Description string     `json:"description"`
	Preparation string     `json:"preparation"`
	Caution     string     `json:"caution"`
	ImageName   string     `json:"image_name"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type DrugDocWithID struct {
	ID string `json:"id"`
	DrugDoc
}

type GetResponseBody struct {
	ID     string  `json:"_id"`
	Source DrugDoc `json:"_source"`
}

type SearchResponseBody struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string  `json:"_id"`
			Source DrugDoc `json:"_source"`
		}
	}
}

type UpdateBody struct {
	Doc DrugDoc `json:"doc"`
}
