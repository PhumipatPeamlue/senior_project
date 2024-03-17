package core

import "context"

type UserServiceInterface interface {
	FindUserByLineUserID(ctx context.Context, lineUserID string) (user User, err error)
	AddNewUser(ctx context.Context, lineUserID string) (err error)
	ChangeTimeSetting(ctx context.Context, lineUserID string, newTimeSetting UserTimeSetting) (err error)
	RemoveUserByLineUserID(ctx context.Context, lineUserID string) (err error)
}

type userService struct {
	repository UserRepositoryInterface
}

// FindUserByLineUserID implements UserServiceInterface.
func (u *userService) FindUserByLineUserID(ctx context.Context, lineUserID string) (user User, err error) {
	user, err = u.repository.ReadByID(ctx, lineUserID)
	return
}

// AddNewUser implements UserServiceInterface.
func (u *userService) AddNewUser(ctx context.Context, lineUserID string) (err error) {
	user := newUser(lineUserID)
	err = u.repository.Create(ctx, user)
	return
}

// ChangeTimeSetting implements UserServiceInterface.
func (u *userService) ChangeTimeSetting(ctx context.Context, lineUserID string, newTimeSetting UserTimeSetting) (err error) {
	user, err := u.FindUserByLineUserID(ctx, lineUserID)
	if err != nil {
		return
	}

	user.changeTimeSetting(newTimeSetting)
	err = u.repository.Update(ctx, user)
	return
}

// RemoveUserByLineUserID implements UserServiceInterface.
func (u *userService) RemoveUserByLineUserID(ctx context.Context, lineUserID string) (err error) {
	err = u.repository.DeleteByID(ctx, lineUserID)
	return
}

func NewUserService(repository UserRepositoryInterface) UserServiceInterface {
	return &userService{
		repository: repository,
	}
}
