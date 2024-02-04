package services

import (
	"context"
	"errors"
	"time"
	"user_service/internal/core/domains"
	"user_service/internal/core/ports"
)

type findByIDResult struct {
	user domains.User
	err  error
}

type userService struct {
	repo ports.UserRepository
}

func (s *userService) GetTimeSettingByID(ctx context.Context, userID string) (uts domains.UserTimeSetting, err error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return
	}

	uts = domains.UserTimeSetting{
		Morning:   user.Morning,
		Noon:      user.Noon,
		Evening:   user.Evening,
		BeforeBed: user.BeforeBed,
	}

	return
}

func (s *userService) CheckUserExists(ctx context.Context, userID string) (exists bool, user domains.User, err error) {
	user, err = s.repo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, NotFoundError) {
			exists = false
			err = nil
		}
		return
	}

	exists = true
	return
}

func (s *userService) AddNewUser(ctx context.Context, userID string) (user domains.User, err error) {
	currentTime := time.Now()
	morning := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 9, 0, 0, 0, currentTime.Location())
	noon := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 12, 0, 0, 0, currentTime.Location())
	evening := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 0, 0, 0, currentTime.Location())
	beforeBed := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 20, 0, 0, 0, currentTime.Location())
	user = domains.User{
		ID:        userID,
		Morning:   morning,
		Noon:      noon,
		Evening:   evening,
		BeforeBed: beforeBed,
	}
	err = s.repo.Save(ctx, user)
	return
}

func (s *userService) ChangeTimeSettings(ctx context.Context, userID string, newUts domains.UserTimeSetting) (err error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return
	}

	user.Morning = newUts.Morning
	user.Noon = newUts.Noon
	user.Evening = newUts.Evening
	user.BeforeBed = newUts.BeforeBed
	err = s.repo.Update(ctx, user)
	return
}

func (s *userService) RemoveUserDataByID(ctx context.Context, userID string) (err error) {
	err = s.repo.DeleteByID(ctx, userID)
	return
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repo: repo,
	}
}
