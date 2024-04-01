package core

import "context"

type iHourNotificationService interface {
	FindHourNotification(ctx context.Context, notificationID string) (hn IHourNotification, err error)
	AddNewHourNotification(ctx context.Context, petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) (err error)
	ChangeHourNotificationInfo(ctx context.Context, notificationID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) (err error)
	RemoveHourNotification(ctx context.Context, notificationID string) (err error)
}

type iPeriodNotificationService interface {
	FindPeriodNotification(ctx context.Context, notificationID string) (pn IPeriodNotification, err error)
	AddNewPeriodNotification(ctx context.Context, petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) (err error)
	ChangePeriodNotificationInfo(ctx context.Context, notificationID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) (err error)
	RemovePeriodNotification(ctx context.Context, notificationID string) (err error)
}

type INotificationService interface {
	FindNotification(ctx context.Context, notificationID string) (notification INotification, err error)
	FindAllPetNotifications(ctx context.Context, petID string) (notifications []INotification, err error)
	RemoveAllPetNotifications(ctx context.Context, petID string) (err error)
	iHourNotificationService
	iPeriodNotificationService
}

type notificationService struct {
	repository                INotificationRepository
	notificationRecordService INotificationRecordService
}

// AddNewHourNotification implements INotificationService.
func (n *notificationService) AddNewHourNotification(ctx context.Context, petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) (err error) {
	hn := newHourNotification(petID, drugInfo, frequencyDayUsage, notifyInfo)
	if err = n.repository.CreateHourNotification(ctx, hn); err != nil {
		return
	}

	err = n.notificationRecordService.AddNewFromHourNotification(ctx, petID, hn.ID(), notifyInfo)
	return
}

// AddNewPeriodNotification implements INotificationService.
func (n *notificationService) AddNewPeriodNotification(ctx context.Context, petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) (err error) {
	pn := newPeriodNotification(petID, drugInfo, frequencyDayUsage, notifyInfo)
	if err = n.repository.CreatePeriodNotification(ctx, pn); err != nil {
		return
	}

	err = n.notificationRecordService.AddNewFromPeriodNotification(ctx, petID, pn.ID(), notifyInfo)
	return
}

// ChangeHourNotificationInfo implements INotificationService.
func (n *notificationService) ChangeHourNotificationInfo(ctx context.Context, notificationID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) (err error) {
	hn, err := n.FindHourNotification(ctx, notificationID)
	if err != nil {
		return
	}

	updateNotification := false
	if hn.NotifyInfo() != notifyInfo {
		updateNotification = true
	}

	hn.changeFrequencyDayUsage(frequencyDayUsage)
	hn.changeDrugInfo(drugInfo)
	hn.changeNotifyInfo(notifyInfo)

	if err = n.repository.UpdateHourNotification(ctx, hn); err != nil {
		return
	}

	if updateNotification {
		err = n.notificationRecordService.ChangeNotifyTimeFromHourNotification(ctx, hn.PetID(), notificationID, notifyInfo)
	}
	return
}

// ChangePeriodNotificationInfo implements INotificationService.
func (n *notificationService) ChangePeriodNotificationInfo(ctx context.Context, notificationID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) (err error) {
	pn, err := n.FindPeriodNotification(ctx, notificationID)
	if err != nil {
		return
	}

	updatedNotification := false
	if pn.NotifyInfo() != notifyInfo {
		pn.changeNotifyInfo(notifyInfo)
		updatedNotification = true
	}

	pn.changeFrequencyDayUsage(frequencyDayUsage)
	pn.changeDrugInfo(drugInfo)

	err = n.repository.UpdatePeriodNotification(ctx, pn)
	if err != nil {
		return
	}

	if updatedNotification {
		err = n.notificationRecordService.ChangeNotifyTimeFromPeriodNotification(ctx, pn.PetID(), notificationID, notifyInfo)
	}

	return
}

// FindAllPetNotifications implements INotificationService.
func (n *notificationService) FindAllPetNotifications(ctx context.Context, petID string) (notifications []INotification, err error) {
	notifications, err = n.repository.ReadByPetID(ctx, petID)
	return
}

// FindHourNotification implements INotificationService.
func (n *notificationService) FindHourNotification(ctx context.Context, notificationID string) (hn IHourNotification, err error) {
	hn, err = n.repository.ReadHourNotificationByID(ctx, notificationID)
	return
}

// FindNotification implements INotificationService.
func (n *notificationService) FindNotification(ctx context.Context, notificationID string) (notification INotification, err error) {
	notification, err = n.repository.ReadByID(ctx, notificationID)
	return
}

// FindPeriodNotification implements INotificationService.
func (n *notificationService) FindPeriodNotification(ctx context.Context, notificationID string) (pn IPeriodNotification, err error) {
	pn, err = n.repository.ReadPeriodNotificationByID(ctx, notificationID)
	return
}

// RemoveAllPetNotifications implements INotificationService.
func (n *notificationService) RemoveAllPetNotifications(ctx context.Context, petID string) (err error) {
	err = n.repository.DeleteByPetID(ctx, petID)
	return
}

// RemoveHourNotification implements INotificationService.
func (n *notificationService) RemoveHourNotification(ctx context.Context, notificationID string) (err error) {
	err = n.repository.DeleteByID(ctx, notificationID)
	return
}

// RemovePeriodNotification implements INotificationService.
func (n *notificationService) RemovePeriodNotification(ctx context.Context, notificationID string) (err error) {
	err = n.repository.DeleteByID(ctx, notificationID)
	return
}

func NewNotificationService(r INotificationRepository, s INotificationRecordService) INotificationService {
	return &notificationService{
		repository:                r,
		notificationRecordService: s,
	}
}
