package repositories

import (
	"context"
	"cronjob/internal/core"
	"cronjob/internal/core/domains"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"time"
)

type notificationRepositorySQL struct {
	db *sql.DB
}

func (n *notificationRepositorySQL) ReadByWaitStatus(ctx context.Context) (notifications []domains.Notification, err error) {
	var query string = `
		SELECT *
		FROM notifications
		WHERE status = 'wait' AND notify_at <= CONVERT_TZ(CURRENT_TIMESTAMP(), 'UTC', 'Asia/Bangkok')
	`
	var id int
	var petID, reminderID, status string
	var notifyAt, createdAt, updatedAt time.Time

	rows, err := n.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&id, &petID, &reminderID, &notifyAt, &status, &createdAt, &updatedAt); err != nil {
			return
		}
		notification := domains.ScanNotification(id, petID, reminderID, status, notifyAt, createdAt, updatedAt)
		notifications = append(notifications, notification)
	}

	return
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

func (n *notificationRepositorySQL) Update(ctx context.Context, notification domains.Notification) (err error) {
	query := `
		UPDATE notifications SET
			status = ?, updated_at = ?
		WHERE id = ?
	`
	id := notification.ID()
	status := notification.Status()
	updatedAt := notification.UpdatedAt()

	_, err = n.db.ExecContext(ctx, query, status, updatedAt, id)
	if errors.Is(err, sql.ErrNoRows) {
		err = core.NewErrNotificationNotFound(err)
	}

	return
}

func NewNotificationRepositorySQL(db *sql.DB) core.NotificationRepositoryInterface {
	return &notificationRepositorySQL{
		db: db,
	}
}
