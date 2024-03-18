package domains

import "time"

type DrugInfo struct {
	DrugName  string `json:"drug_name"`
	DrugUsage string `json:"drug_usage"`
}

type HourNotifyInfo struct {
	FirstUsage *time.Time `json:"first_usage"`
	Every      *int       `json:"every"`
}

type PeriodNotifyInfo struct {
	Morning   *time.Time `json:"morning"`
	Noon      *time.Time `json:"noon"`
	Evening   *time.Time `json:"evening"`
	BeforeBed *time.Time `json:"before_bed"`
}

type Reminder struct {
	id                string
	petID             string
	reminderType      string
	drugInfo          DrugInfo
	frequencyDayUsage int
	renewIn           int
	hourNotifyInfo    HourNotifyInfo
	periodNotifyInfo  PeriodNotifyInfo
	createdAt         time.Time
	updatedAt         time.Time
}

func (r *Reminder) ID() string {
	return r.id
}

func (r *Reminder) PetID() string {
	return r.petID
}

func (r *Reminder) Type() string {
	return r.reminderType
}

func (r *Reminder) FrequencyDayUsage() int {
	return r.frequencyDayUsage
}

func (r *Reminder) RenewIn() int {
	return r.renewIn
}

func (r *Reminder) DrugInfo() DrugInfo {
	return r.drugInfo
}

func (r *Reminder) HourNotifyInfo() HourNotifyInfo {
	return r.hourNotifyInfo
}

func (r *Reminder) PeriodNotifyInfo() PeriodNotifyInfo {
	return r.periodNotifyInfo
}

func (r *Reminder) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *Reminder) ReNew() {
	r.renewIn = r.frequencyDayUsage
	r.updatedAt = time.Now().Local()
}

func (r *Reminder) DecrementRenew() {
	r.renewIn -= 1
}

func ScanReminder(id, petID, reminderType string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, hourNotifyInfo HourNotifyInfo, periodNotifyInfo PeriodNotifyInfo, createdAt, updatedAt time.Time) Reminder {
	return Reminder{
		id:                id,
		petID:             petID,
		reminderType:      reminderType,
		drugInfo:          drugInfo,
		frequencyDayUsage: frequencyDayUsage,
		renewIn:           renewIn,
		hourNotifyInfo:    hourNotifyInfo,
		periodNotifyInfo:  periodNotifyInfo,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}
