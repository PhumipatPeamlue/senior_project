package repositories

import (
	"context"
	"database/sql"
	"errors"
	"user_web_service/internal/core"

	"github.com/go-sql-driver/mysql"
)

type notificationRecordRepositorySQL struct {
	db *sql.DB
}

func (n *notificationRecordRepositorySQL) Create(ctx context.Context, nr core.INotificationRecord) (err error) {
	query := `
		INSERT INTO notification_records (
			pet_id, notification_id, notify_at, status, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	petID := nr.PetID()
	notificationID := nr.NotificationID()
	notifyAt := nr.NotifyAt()
	status := nr.Status()
	createdAt, updatedAt := nr.CreatedAt(), nr.UpdatedAt()

	_, err = n.db.ExecContext(ctx, query, petID, notificationID, notifyAt, status, createdAt, updatedAt)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = &core.ErrDuplicate{OriginalError: err}
		}
	}

	return
}

func (n *notificationRecordRepositorySQL) DeleteTodayAndWaitStatusByNotificationID(ctx context.Context, notificationID string) (err error) {
	query := `
		DELETE FROM notification_records
		WHERE notification_id = ? AND status = 'wait' AND DATE(notify_at) = CURDATE()
	`
	_, err = n.db.ExecContext(ctx, query, notificationID)
	return
}

func NewNotificationRecordRepositorySQL(db *sql.DB) core.INotificationRecordRepository {
	return &notificationRecordRepositorySQL{
		db: db,
	}
}
