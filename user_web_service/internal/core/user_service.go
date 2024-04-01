package core

import "context"

type IUserService interface {
	FindUserByLineUserID(ctx context.Context, lineUserID string) (user IUser, err error)
	AddNewUser(ctx context.Context, lineUserID string) (err error)
	ChangeTimeSetting(ctx context.Context, lineUserID string, newTimeSetting IUserTimeSetting) (err error)
	RemoveUserByLineUserID(ctx context.Context, lineUserID string) (err error)
}

type userService struct {
	repository IUserRepository
}

// AddNewUser implements IUserService.
func (u *userService) AddNewUser(ctx context.Context, lineUserID string) (err error) {
	user := newUser(lineUserID)
	err = u.repository.Create(ctx, user)
	return
}

// ChangeTimeSetting implements IUserService.
func (u *userService) ChangeTimeSetting(ctx context.Context, lineUserID string, newTimeSetting IUserTimeSetting) (err error) {
	user, err := u.FindUserByLineUserID(ctx, lineUserID)
	if err != nil {
		return
	}

	user.changeTimeSetting(newTimeSetting)
	err = u.repository.Update(ctx, user)
	return
}

// FindUserByLineUserID implements IUserService.
func (u *userService) FindUserByLineUserID(ctx context.Context, lineUserID string) (user IUser, err error) {
	user, err = u.repository.ReadByID(ctx, lineUserID)
	return
}

// RemoveUserByLineUserID implements IUserService.
func (u *userService) RemoveUserByLineUserID(ctx context.Context, lineUserID string) (err error) {
	err = u.repository.DeleteByID(ctx, lineUserID)
	return
}

func NewUserService(r IUserRepository) IUserService {
	return &userService{
		repository: r,
	}
}
