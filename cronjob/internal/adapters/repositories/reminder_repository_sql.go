package repositories

import (
	"context"
	"cronjob/internal/core"
	"cronjob/internal/core/domains"
	"database/sql"
	"errors"
	"time"
)

type reminderRepositorySQL struct {
	db *sql.DB
}

func (r *reminderRepositorySQL) ReadAll(ctx context.Context) (reminders []domains.Reminder, err error) {
	query := `
		SELECT id, pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type
		FROM reminders
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, petID, reminderType string
		var drugInfo domains.DrugInfo
		var frequencyDayUsage, renewIn int
		var createdAt, updatedAt time.Time
		var hourNotifyInfo domains.HourNotifyInfo
		var periodNotifyInfo domains.PeriodNotifyInfo
		err = rows.Scan(&id, &petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &reminderType)
		if err != nil {
			return
		}
		reminder := domains.ScanReminder(id, petID, reminderType, drugInfo, frequencyDayUsage, renewIn, hourNotifyInfo, periodNotifyInfo, createdAt, updatedAt)
		reminders = append(reminders, reminder)
	}

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
	var hourNotifyInfo domains.HourNotifyInfo
	var periodNotifyInfo domains.PeriodNotifyInfo

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &reminderType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrReminderNotFound(err)
		}

		return
	}

	reminder = domains.ScanReminder(id, petID, reminderType, drugInfo, frequencyDayUsage, renewIn, hourNotifyInfo, periodNotifyInfo, createdAt, updatedAt)
	return
}

func (r *reminderRepositorySQL) ReadAllZeroRenew(ctx context.Context) (reminders []domains.Reminder, err error) {
	query := `
		SELECT
		    r.id AS id,
		    r.pet_id AS pet_id,
		    r.type AS type,
		    r.drug_name AS drug_name,
		    r.drug_usage AS drug_usage,
		    r.frequency_day_usage AS frequency_day_usage,
			r.renew_in AS renew_in,
			r.created_at AS created_at,
			r.updated_at AS updated_at,
			hi.first_usage AS first_usage,
			hi.every AS every,
			pi.morning AS morning,
			pi.noon AS noon,
			pi.evening AS evening,
			pi.before_bed AS before_bed
		FROM
		    reminders AS r
		LEFT JOIN hour_info AS hi on r.id = hi.reminder_id
		LEFT JOIN period_info AS pi on r.id = pi.reminder_id
		WHERE
		    renew_in = 0
	`
	var id, petID, reminderType string
	var drugInfo domains.DrugInfo
	var frequencyDayUsage, renewIn int
	var hourNotifyInfo domains.HourNotifyInfo
	var periodNotifyInfo domains.PeriodNotifyInfo
	var createdAt, updatedAt time.Time

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &petID, &reminderType, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn,
			&createdAt, &updatedAt, &hourNotifyInfo.FirstUsage, &hourNotifyInfo.Every, &periodNotifyInfo.Morning,
			&periodNotifyInfo.Noon, &periodNotifyInfo.Evening, &periodNotifyInfo.BeforeBed)
		if err != nil {
			return
		}
		reminder := domains.ScanReminder(id, petID, reminderType, drugInfo, frequencyDayUsage, renewIn, hourNotifyInfo, periodNotifyInfo, createdAt, updatedAt)
		reminders = append(reminders, reminder)
	}

	return
}

func (r *reminderRepositorySQL) UpdateRenew(ctx context.Context, reminder domains.Reminder) (err error) {
	query := "UPDATE reminders SET renew_in = ?, updated_at = ? WHERE id = ?"
	_, err = r.db.ExecContext(ctx, query, reminder.RenewIn(), reminder.UpdatedAt(), reminder.ID())
	return
}

func NewReminderRepositorySQL(db *sql.DB) core.ReminderRepositoryInterface {
	return &reminderRepositorySQL{
		db: db,
	}
}
