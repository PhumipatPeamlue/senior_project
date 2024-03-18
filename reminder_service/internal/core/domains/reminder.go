package domains

import (
	"time"

	"github.com/google/uuid"
)

type DrugInfo struct {
	DrugName  string `json:"drug_name"`
	DrugUsage string `json:"drug_usage"`
}

type HourNotifyInfo struct {
	FirstUsage time.Time `json:"first_usage"`
	Every      int       `json:"every"`
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

func (r *Reminder) CreatedAt() time.Time {
	return r.createdAt
}

func (r *Reminder) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *Reminder) ChangeFrequencyDayUsage(newFrequencyDayUsage int) {
	r.frequencyDayUsage = newFrequencyDayUsage
	r.renewIn = 0
	r.updatedAt = time.Now().Local()
}

func (r *Reminder) ChangeDrugInfo(updatedDrugInfo DrugInfo) {
	r.drugInfo = updatedDrugInfo
	r.updatedAt = time.Now().Local()
}

func ScanReminder(id, petID, reminderType string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time) Reminder {
	return Reminder{
		id:                id,
		petID:             petID,
		reminderType:      reminderType,
		frequencyDayUsage: frequencyDayUsage,
		renewIn:           renewIn,
		drugInfo:          drugInfo,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

type HourReminder struct {
	Reminder
	notifyInfo HourNotifyInfo
}

func (h *HourReminder) NotifyInfo() HourNotifyInfo {
	return h.notifyInfo
}

func (h *HourReminder) ChangeNotifyInfo(updatedNotifyInfo HourNotifyInfo) {
	h.notifyInfo = updatedNotifyInfo
	h.updatedAt = time.Now().Local()
}

func ScanHourReminder(id, petID string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time, notifyInfo HourNotifyInfo) HourReminder {
	reminder := ScanReminder(id, petID, "hour", drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
	return HourReminder{
		Reminder:   reminder,
		notifyInfo: notifyInfo,
	}
}

func NewHourReminder(petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) HourReminder {
	id := uuid.New().String()
	now := time.Now().Local()
	hr := ScanHourReminder(id, petID, drugInfo, frequencyDayUsage, 0, now, now, notifyInfo)
	return hr
}

type PeriodReminder struct {
	Reminder
	notifyInfo PeriodNotifyInfo
}

func (p *PeriodReminder) NotifyInfo() PeriodNotifyInfo {
	return p.notifyInfo
}

func (p *PeriodReminder) ChangeNotifyInfo(updatedNotifyInfo PeriodNotifyInfo) {
	p.notifyInfo = updatedNotifyInfo
	p.updatedAt = time.Now().Local()
}

func ScanPeriodReminder(id, petID string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time, notifyInfo PeriodNotifyInfo) PeriodReminder {
	reminder := ScanReminder(id, petID, "period", drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
	return PeriodReminder{
		Reminder:   reminder,
		notifyInfo: notifyInfo,
	}
}

func NewPeriodReminder(petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) PeriodReminder {
	id := uuid.New().String()
	now := time.Now().Local()
	pr := ScanPeriodReminder(id, petID, drugInfo, frequencyDayUsage, 0, now, now, notifyInfo)
	return pr
}
