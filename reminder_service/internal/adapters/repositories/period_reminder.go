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

type periodReminderRepositoryMySQL struct {
	reminderRepo ports.ReminderRepository
	db           *sql.DB
}

func (r *periodReminderRepositoryMySQL) FindByID(ctx context.Context, id string) (pr domains.PeriodReminder, err error) {
	query := "SELECT reminders.id, reminders.pet_id, reminders.type, reminders.drug_name, reminders.drug_usage, reminders.frequency, period_reminder_info.morning, period_reminder_info.noon, period_reminder_info.evening, period_reminder_info.before_bed FROM reminders INNER JOIN period_reminder_info ON reminders.id = period_reminder_info.reminder_id WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&pr.ID, &pr.PetID, &pr.Type, &pr.DrugName, &pr.DrugUsage, &pr.Frequency, &pr.Morning, &pr.Noon, &pr.Evening, &pr.BeforeBed); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = services.NotFoundError
		}
	}
	return
}

func (r *periodReminderRepositoryMySQL) Save(ctx context.Context, pr domains.PeriodReminder) (err error) {
	if err = r.reminderRepo.Save(ctx, pr.Reminder); err != nil {
		return
	}

	query := "INSERT INTO period_reminder_info (reminder_id, morning, noon, evening, before_bed) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, pr.ID, pr.Morning, pr.Noon, pr.Evening, pr.BeforeBed); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func (r *periodReminderRepositoryMySQL) Update(ctx context.Context, pr domains.PeriodReminder) (err error) {
	if err = r.reminderRepo.Update(ctx, pr.Reminder); err != nil {
		return
	}

	query := "UPDATE period_reminder_info SET morning = ?, noon = ?, evening = ?, before_bed = ? WHERE reminder_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, pr.Morning, pr.Noon, pr.Evening, pr.BeforeBed, pr.ID)
	return
}

func NewPeriodReminderRepositoryMySQL(reminderRepo ports.ReminderRepository, db *sql.DB) ports.PeriodReminderRepository {
	return &periodReminderRepositoryMySQL{
		reminderRepo: reminderRepo,
		db:           db,
	}
}
