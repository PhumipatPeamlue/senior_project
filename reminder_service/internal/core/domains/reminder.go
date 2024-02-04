package domains

type Reminder struct {
	ID        string `json:"reminder_id"`
	PetID     string `json:"pet_id"`
	Type      string `json:"reminder_type"`
	DrugName  string `json:"drug_name"`
	DrugUsage string `json:"drug_usage"`
	Frequency string `json:"frequency"`
}
