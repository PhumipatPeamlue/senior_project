package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"notification_service/internal/core"
	"notification_service/internal/core/domains"
)

type notificationRepositorySQL struct {
	db *sql.DB
}

func (n *notificationRepositorySQL) Create(ctx context.Context, notification domains.Notification) (err error) {
	query := `
		INSERT INTO notifications (
			pet_id, reminder_id, notify_at, status, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	petID := notification.PetID()
	reminderID := notification.ReminderID()
	notifyAt := notification.NotifyAt()
	status := notification.Status()
	createdAt, updatedAt := notification.CreatedAt(), notification.UpdatedAt()

	_, err = n.db.ExecContext(ctx, query, petID, reminderID, notifyAt, status, createdAt, updatedAt)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = core.NewErrNotificationDuplicate(err)
		}
	}

	return
}

func (n *notificationRepositorySQL) DeleteTodayAndWaitStatusByReminderID(ctx context.Context, reminderID string) (err error) {
	query := `
		DELETE FROM notifications
		WHERE reminder_id = ? AND status = 'wait' AND DATE(notify_at) = CURDATE()
	`
	_, err = n.db.ExecContext(ctx, query, reminderID)
	return
}

func NewNotificationSQL(db *sql.DB) core.NotificationRepositoryInterface {
	return &notificationRepositorySQL{
		db: db,
	}
}
