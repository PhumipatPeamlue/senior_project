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

type notificationRepositoryMySQL struct {
	db *sql.DB
}

func (r *notificationRepositoryMySQL) DeleteByReminderID(ctx context.Context, reminderID string) (err error) {
	query := "DELETE FROM notifications WHERE reminder_id = ? AND status = 'not sent'"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reminderID)
	return
}

func (r *notificationRepositoryMySQL) Save(ctx context.Context, notification domains.Notification) (err error) {
	query := "INSERT INTO notifications (reminder_id, user_id, time, status) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, notification.ReminderID, notification.UserID, notification.Time, notification.Status); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func NewNotificationRepositoryMySQL(db *sql.DB) ports.NotificationRepository {
	return &notificationRepositoryMySQL{
		db: db,
	}
}
