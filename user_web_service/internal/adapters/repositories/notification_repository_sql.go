package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"user_web_service/internal/core"
)

type notificationRepositorySQL struct {
	db *sql.DB
}

func (r *notificationRepositorySQL) ReadHourNotificationByID(ctx context.Context, id string) (hn core.IHourNotification, err error) {
	var query string = `
		SELECT
			n.pet_id AS pet_id, 
			n.drug_name AS drug_name, 
			n.drug_usage AS drug_usage, 
			n.frequency_day_usage AS frequency_day_usage,
			n.renew_in AS renew_in,
			h.first_usage AS first_usage, 
			h.every AS every,
			n.created_at AS created_at, 
			n.updated_at AS updated_at
		FROM 
			notifications AS n
		INNER JOIN hour_info AS h ON 
			n.id = h.notification_id
		WHERE
			n.id = ?
	`
	var petID string
	var drugInfo core.DrugInfo
	var frequencyDayUsage, renewIn int
	var notifyInfo core.HourNotifyInfo
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &notifyInfo.FirstUsage, &notifyInfo.Every, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = &core.ErrNotFound{OriginalError: err}
		}

		return
	}

	hn = core.ScanHourNotification(id, petID, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt, notifyInfo)
	return
}

func (r *notificationRepositorySQL) CreateHourNotification(ctx context.Context, hn core.IHourNotification) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// INSERT notification
	query1 := `
		INSERT INTO notifications (id, pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	id := hn.ID()
	petID := hn.PetID()
	drugInfo := hn.DrugInfo()
	frequencyDayUsage := hn.FrequencyDayUsage()
	renewIn := hn.RenewIn()
	createdAt, updatedAt := hn.CreatedAt(), hn.UpdatedAt()
	notificationType := hn.Type()
	drugName, drugUsage := drugInfo.DrugName, drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, id, petID, drugName, drugUsage, frequencyDayUsage, renewIn, createdAt, updatedAt, notificationType)
	if err != nil {
		tx.Rollback()
		return
	}

	// INSERT hour_info
	query2 := `
		INSERT INTO hour_info (notification_id, first_usage, every)
		VALUES (?, ?, ?)
	`
	notifyInfo := hn.NotifyInfo()
	firstUsage, every := notifyInfo.FirstUsage, notifyInfo.Every
	_, err = tx.ExecContext(ctx, query2, id, firstUsage, every)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *notificationRepositorySQL) UpdateHourNotification(ctx context.Context, hn core.IHourNotification) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// UPDATE notifications statement
	query1 := `
		UPDATE notifications SET drug_name = ?, drug_usage = ?, updated_at = ?, frequency_day_usage = ?, renew_in = ?
		WHERE id = ?
	`
	id := hn.ID()
	drugInfo := hn.DrugInfo()
	frequencyDayUsage := hn.FrequencyDayUsage()
	renewIn := hn.RenewIn()
	updatedAt := hn.UpdatedAt()
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
		WHERE notification_id = ?
	`
	notifyInfo := hn.NotifyInfo()
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

func (r *notificationRepositorySQL) ReadPeriodNotificationByID(ctx context.Context, id string) (pn core.IPeriodNotification, err error) {
	var query string = `
		SELECT
			n.pet_id AS pet_id, 
			n.drug_name AS drug_name, 
			n.drug_usage AS drug_usage, 
			n.frequency_day_usage AS frequency_day_usage, 
			n.renew_in AS renew_in,
			p.morning AS morning,
			p.noon AS noon,
			p.evening AS evening,
			p.before_bed AS before_bed,
			n.created_at AS created_at, 
			n.updated_at AS updated_at
		FROM 
			notifications AS n
		INNER JOIN period_info AS p ON 
			n.id = p.notification_id
		WHERE
			n.id = ?
	`
	var petID string
	var drugInfo core.DrugInfo
	var frequencyDayUsage, renewIn int
	var notifyInfo core.PeriodNotifyInfo
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &notifyInfo.Morning,
		&notifyInfo.Noon, &notifyInfo.Evening, &notifyInfo.BeforeBed, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = &core.ErrNotFound{OriginalError: err}
		}

		return
	}

	pn = core.ScanPeriodNotification(id, petID, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt, notifyInfo)
	return
}

func (r *notificationRepositorySQL) CreatePeriodNotification(ctx context.Context, pn core.IPeriodNotification) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// INSERT notifications
	query1 := `
		INSERT INTO notifications (id, pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	id := pn.ID()
	petID := pn.PetID()
	drugInfo := pn.DrugInfo()
	frequencyDayUsage := pn.FrequencyDayUsage()
	renewIn := pn.RenewIn()
	createdAt, updatedAt := pn.CreatedAt(), pn.UpdatedAt()
	notificationType := pn.Type()
	drugName, drugUsage := drugInfo.DrugName, drugInfo.DrugUsage

	_, err = tx.ExecContext(ctx, query1, id, petID, drugName, drugUsage, frequencyDayUsage, renewIn, createdAt, updatedAt, notificationType)
	if err != nil {
		tx.Rollback()
		return
	}

	// INSERT period_info
	query2 := `
		INSERT INTO period_info (notification_id, morning, noon, evening, before_bed)
		VALUES (?, ?, ?, ?, ?)
	`
	notificationID := pn.ID()
	notifyInfo := pn.NotifyInfo()
	morning, noon, evening, beforeBed := notifyInfo.Morning, notifyInfo.Noon, notifyInfo.Evening, notifyInfo.BeforeBed
	_, err = tx.ExecContext(ctx, query2, notificationID, morning, noon, evening, beforeBed)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *notificationRepositorySQL) UpdatePeriodNotification(ctx context.Context, pn core.IPeriodNotification) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// UPDATE notifications statement
	query1 := `
		UPDATE notifications SET drug_name = ?, drug_usage = ?, frequency_day_usage = ?, renew_in = ?, updated_at = ?
		WHERE id = ?
	`
	id := pn.ID()
	drugInfo := pn.DrugInfo()
	frequencyDayUsage := pn.FrequencyDayUsage()
	renewIn := pn.RenewIn()
	updatedAt := pn.UpdatedAt()
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
		WHERE notification_id = ?
	`
	notifyInfo := pn.NotifyInfo()
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

func (r *notificationRepositorySQL) ReadByID(ctx context.Context, id string) (notification core.INotification, err error) {
	query := `
		SELECT pet_id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type
		FROM notifications
		WHERE id = ?
	`
	var petID, notificationType string
	var drugInfo core.DrugInfo
	var frequencyDayUsage, renewIn int
	var createdAt, updatedAt time.Time

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&petID, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &notificationType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = &core.ErrNotFound{OriginalError: err}
		}

		return
	}

	notification = core.ScanNotification(id, petID, notificationType, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
	return
}

func (r *notificationRepositorySQL) ReadByPetID(ctx context.Context, petID string) (notifications []core.INotification, err error) {
	query := `
		SELECT id, drug_name, drug_usage, frequency_day_usage, renew_in, created_at, updated_at, type
		FROM notifications
		WHERE pet_id = ?
	`
	var id, notificationType string
	var drugInfo core.DrugInfo
	var frequencyDayUsage, renewIn int
	var createdAt, updatedAt time.Time

	rows, err := r.db.QueryContext(ctx, query, petID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &drugInfo.DrugName, &drugInfo.DrugUsage, &frequencyDayUsage, &renewIn, &createdAt, &updatedAt, &notificationType)
		if err != nil {
			return
		}
		notification := core.ScanNotification(id, petID, notificationType, drugInfo, frequencyDayUsage, renewIn, createdAt, updatedAt)
		notifications = append(notifications, notification)
	}

	if len(notifications) == 0 {
		err = &core.ErrNotFound{OriginalError: sql.ErrNoRows}
	}

	return
}

func (r *notificationRepositorySQL) DeleteByPetID(ctx context.Context, petID string) (err error) {
	query := "DELETE FROM notifications WHERE pet_id = ?"
	_, err = r.db.ExecContext(ctx, query, petID)
	return
}

func (r *notificationRepositorySQL) DeleteByID(ctx context.Context, id string) (err error) {
	query := "DELETE FROM notifications WHERE id = ?"
	_, err = r.db.ExecContext(ctx, query, id)
	return
}

func NewNotificationRepositorySQL(db *sql.DB) core.INotificationRepository {
	return &notificationRepositorySQL{
		db: db,
	}
}
