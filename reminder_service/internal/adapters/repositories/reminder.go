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

type reminderRepositoryMySQL struct {
	db *sql.DB
}

func (r *reminderRepositoryMySQL) FindByPetID(ctx context.Context, petID string) (reminders []domains.Reminder, err error) {
	query := "SELECT id, type, drug_name, drug_usage, frequency FROM reminders WHERE pet_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, petID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, rt, name, usage, f string
		if err = rows.Scan(&id, &rt, &name, &usage, &f); err != nil {
			return
		}
		reminder := domains.Reminder{
			ID:        id,
			PetID:     petID,
			Type:      rt,
			DrugName:  name,
			DrugUsage: usage,
			Frequency: f,
		}
		reminders = append(reminders, reminder)
	}

	if len(reminders) == 0 {
		err = services.NotFoundError
	}
	return
}

func (r *reminderRepositoryMySQL) Save(ctx context.Context, reminder domains.Reminder) (err error) {
	query := "INSERT INTO reminders (id, pet_id, type, drug_name, drug_usage, frequency) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, reminder.ID, reminder.PetID, reminder.Type, reminder.DrugName, reminder.DrugUsage, reminder.Frequency); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func (r *reminderRepositoryMySQL) Update(ctx context.Context, reminder domains.Reminder) (err error) {
	query := "UPDATE reminders SET drug_name = ?, drug_usage = ?, frequency = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reminder.DrugName, reminder.DrugUsage, reminder.Frequency, reminder.ID)
	return
}

func (r *reminderRepositoryMySQL) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM reminders WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = services.NotFoundError
	}
	return
}

func NewReminderRepositoryMySQL(db *sql.DB) ports.ReminderRepository {
	return &reminderRepositoryMySQL{
		db: db,
	}
}
