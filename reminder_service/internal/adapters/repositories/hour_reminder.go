package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"reminder_service/internal/core/domains"
	"reminder_service/internal/core/ports"
	"reminder_service/internal/core/services"
)

type hourReminderRepositoryMySQL struct {
	reminderRepo ports.ReminderRepository
	db           *sql.DB
}

func (r *hourReminderRepositoryMySQL) FindByID(ctx context.Context, id string) (hr domains.HourReminder, err error) {
	query := "SELECT reminders.id, reminders.pet_id, reminders.type, reminders.drug_name, reminders.drug_usage, reminders.frequency, hour_reminder_info.first_usage, hour_reminder_info.every FROM reminders INNER JOIN hour_reminder_info ON reminders.id = hour_reminder_info.reminder_id WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&hr.ID, &hr.PetID, &hr.Type, &hr.DrugName, &hr.DrugUsage, &hr.Frequency, &hr.FirstUsage, &hr.Every); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = services.NotFoundError
		}
	}
	return
}

func (r *hourReminderRepositoryMySQL) Save(ctx context.Context, hr domains.HourReminder) (err error) {
	if err = r.reminderRepo.Save(ctx, hr.Reminder); err != nil {
		return
	}

	query := "INSERT INTO hour_reminder_info (reminder_id, first_usage, every) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, hr.ID, hr.FirstUsage, hr.Every); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func (r *hourReminderRepositoryMySQL) Update(ctx context.Context, hr domains.HourReminder) (err error) {
	if err = r.reminderRepo.Update(ctx, hr.Reminder); err != nil {
		return
	}

	query := "UPDATE hour_reminder_info SET first_usage = ?, every = ? WHERE reminder_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, hr.FirstUsage, hr.Every, hr.ID)
	return
}

func NewHourReminderRepositoryMySQL(reminderRepo ports.ReminderRepository, db *sql.DB) ports.HourReminderRepository {
	return &hourReminderRepositoryMySQL{
		reminderRepo: reminderRepo,
		db:           db,
	}
}
