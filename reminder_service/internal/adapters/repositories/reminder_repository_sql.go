package repositories

import (
	"context"
	"database/sql"
	"errors"
	"reminder_service/internal/core"
	"reminder_service/internal/core/domains"
	"time"
)

type reminderRepositorySQL struct {
	db *sql.DB
}

func (r *reminderRepositorySQL) ReadHourReminderByID(ctx context.Context, id string) (hr domains.HourReminder, err error) {
	var query string = `
		SELECT
			r.pet_id AS pet_id, 
			r.drug_name AS drug_name, 
			r.drug_usage AS drug_usage, 
			r.frequency_day_usage AS frequency_day_usage,
			r.renew_in AS renew_in,
			h.first_usage AS first_usage, 
			h.every AS every,
			r.created_at AS created_at, 
			r.updated_at AS updated_at
		FROM 
			reminders AS r
		INNER JOIN hour_info AS h ON 
			r.id = h.reminder_id
		WHERE
			r.id = ?
	`
	var petID string
	var drugInfo domains.DrugInfo
	var frequencyDayUsage, renewIn int
	var notifyInfo domains.HourNotifyInfo
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &notifyInfo.FirstUsage, &notifyInfo.Every, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrReminderNotFound(err)
		}

		return
	}

	hr = domains.ScanHourReminder(id, petID, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt, notifyInfo)
	return
}

func (r *reminderRepositorySQL) CreateHourReminder(ctx context.Context, hr domains.HourReminder) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// INSERT reminders
	query1 := `
		INSERT INTO reminders (id, pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	id := hr.ID()
	petID := hr.PetID()
	drugInfo := hr.DrugInfo()
	frequencyDayUsage := hr.FrequencyDayUsage()
	renewIn := hr.RenewIn()
	createdAt, updatedAt := hr.CreatedAt(), hr.UpdatedAt()
	reminderType := hr.Type()
	drugName, drugUsage := drugInfo.DrugName, drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, id, petID, drugName, drugUsage, frequencyDayUsage, renewIn, createdAt, updatedAt, reminderType)
	if err != nil {
		tx.Rollback()
		return
	}

	// INSERT hour_info
	query2 := `
		INSERT INTO hour_info (reminder_id, first_usage, every)
		VALUES (?, ?, ?)
	`
	notifyInfo := hr.NotifyInfo()
	firstUsage, every := notifyInfo.FirstUsage, notifyInfo.Every
	_, err = tx.ExecContext(ctx, query2, id, firstUsage, every)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *reminderRepositorySQL) UpdateHourReminder(ctx context.Context, hr domains.HourReminder) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// UPDATE reminders statement
	query1 := `
		UPDATE reminders SET drug_name = ?, drug_usage = ?, updated_at = ?, frequency_day_usage = ?, renew_in = ?
		WHERE id = ?
	`
	id := hr.ID()
	drugInfo := hr.DrugInfo()
	frequencyDayUsage := hr.FrequencyDayUsage()
	renewIn := hr.RenewIn()
	updatedAt := hr.UpdatedAt()
	drugName := drugInfo.DrugName
	drugUsage := drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, drugName, drugUsage, updatedAt, frequencyDayUsage, renewIn, id)
	if err != nil {
		tx.Rollback()
		return
	}

	// UPDATE hour_info statement
	query2 := `
		UPDATE hour_info SET first_usage = ?, every = ?
		WHERE reminder_id = ?
	`
	notifyInfo := hr.NotifyInfo()
	firstUsage := notifyInfo.FirstUsage
	every := notifyInfo.Every
	_, err = tx.ExecContext(ctx, query2, firstUsage, every, id)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *reminderRepositorySQL) ReadPeriodReminderByID(ctx context.Context, id string) (pr domains.PeriodReminder, err error) {
	var query string = `
		SELECT
			r.pet_id AS pet_id, 
			r.drug_name AS drug_name, 
			r.drug_usage AS drug_usage, 
			r.frequency_day_usage AS frequency_day_usage, 
			r.renew_in AS renew_in,
			p.morning AS morning,
			p.noon AS noon,
			p.evening AS evening,
			p.before_bed AS before_bed,
			r.created_at AS created_at, 
			r.updated_at AS updated_at
		FROM 
			reminders AS r
		INNER JOIN period_info AS p ON 
			r.id = p.reminder_id
		WHERE
			r.id = ?
	`
	var petID string
	var drugInfo domains.DrugInfo
	var frequencyDayUsage, renewIn int
	var notifyInfo domains.PeriodNotifyInfo
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &notifyInfo.Morning,
		&notifyInfo.Noon, &notifyInfo.Evening, &notifyInfo.BeforeBed, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrReminderNotFound(err)
		}

		return
	}

	pr = domains.ScanPeriodReminder(id, petID, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt, notifyInfo)
	return
}

func (r *reminderRepositorySQL) CreatePeriodReminder(ctx context.Context, pr domains.PeriodReminder) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// INSERT reminders
	query1 := `
		INSERT INTO reminders (id, pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	id := pr.ID()
	petID := pr.PetID()
	drugInfo := pr.DrugInfo()
	frequencyDayUsage := pr.FrequencyDayUsage()
	renewIn := pr.RenewIn()
	createdAt, updatedAt := pr.CreatedAt(), pr.UpdatedAt()
	reminderType := pr.Type()
	drugName, drugUsage := drugInfo.DrugName, drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, id, petID, drugName, drugUsage, frequencyDayUsage, renewIn, createdAt, updatedAt, reminderType)
	if err != nil {
		tx.Rollback()
		return
	}

	// INSERT period_info
	query2 := `
		INSERT INTO period_info (reminder_id, morning, noon, evening, before_bed)
		VALUES (?, ?, ?, ?, ?)
	`
	reminderID := pr.ID()
	notifyInfo := pr.NotifyInfo()
	morning, noon, evening, beforeBed := notifyInfo.Morning, notifyInfo.Noon, notifyInfo.Evening, notifyInfo.BeforeBed
	_, err = tx.ExecContext(ctx, query2, reminderID, morning, noon, evening, beforeBed)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *reminderRepositorySQL) UpdatePeriodReminder(ctx context.Context, pr domains.PeriodReminder) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// UPDATE reminders statement
	query1 := `
		UPDATE reminders SET drug_name = ?, drug_usage = ?, frequency_day_usage = ?, renew_in = ?, updated_at = ?
		WHERE id = ?
	`
	id := pr.ID()
	drugInfo := pr.DrugInfo()
	frequencyDayUsage := pr.FrequencyDayUsage()
	renewIn := pr.RenewIn()
	updatedAt := pr.UpdatedAt()
	drugName := drugInfo.DrugName
	drugUsage := drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, drugName, drugUsage, frequencyDayUsage, renewIn, updatedAt, id)
	if err != nil {
		tx.Rollback()
		return
	}

	// UPDATE period_info statement
	query2 := `
		UPDATE period_info SET morning = ?, noon = ?, evening = ?, before_bed = ?
		WHERE reminder_id = ?
	`
	notifyInfo := pr.NotifyInfo()
	morning := notifyInfo.Morning
	noon := notifyInfo.Noon
	evening := notifyInfo.Evening
	beforeBed := notifyInfo.BeforeBed
	_, err = tx.ExecContext(ctx, query2, morning, noon, evening, beforeBed, id)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *reminderRepositorySQL) ReadByID(ctx context.Context, id string) (reminder domains.Reminder, err error) {
	query := `
		SELECT pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type
		FROM reminders
		WHERE id = ?
	`
	var petID, reminderType string
	var drugInfo domains.DrugInfo
	var frequencyDayUsage, renewIn int
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &reminderType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrReminderNotFound(err)
		}

		return
	}

	reminder = domains.ScanReminder(id, petID, reminderType, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
	return
}

func (r *reminderRepositorySQL) ReadByPetID(ctx context.Context, petID string) (reminders []domains.Reminder, err error) {
	query := `
		SELECT id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type
		FROM reminders
		WHERE pet_id = ?
	`
	var id, reminderType string
	var drugInfo domains.DrugInfo
	var frequencyDayUsage, renewIn int
	var createdAt, updatedAt time.Time

	rows, err := r.db.QueryContext(ctx, query, petID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &reminderType)
		if err != nil {
			return
		}
		reminder := domains.ScanReminder(id, petID, reminderType, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
		reminders = append(reminders, reminder)
	}

	if len(reminders) == 0 {
		err = core.NewErrReminderNotFound(sql.ErrNoRows)
	}

	return
}

func (r *reminderRepositorySQL) DeleteByPetID(ctx context.Context, petID string) (err error) {
	query := "DELETE FROM reminders WHERE pet_id = ?"
	_, err = r.db.ExecContext(ctx, query, petID)
	return
}

func (r *reminderRepositorySQL) DeleteByID(ctx context.Context, id string) (err error) {
	query := "DELETE FROM reminders WHERE id = ?"
	_, err = r.db.ExecContext(ctx, query, id)
	return
}

func NewReminderRepositorySQL(db *sql.DB) core.ReminderRepositoryInterface {
	return &reminderRepositorySQL{
		db: db,
	}
}
