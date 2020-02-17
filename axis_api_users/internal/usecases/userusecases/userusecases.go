package userusecases

import (
	"context"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	userrepo "github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/repositories/userrepositories"
	userservice "github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/services/userservices"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apierrors"
)

//UserUseCase ...
type UserUseCase interface {
	GetByEmail(context.Context, string) (*users.User, apierrors.APIErr)
	Save(context.Context, *users.User) apierrors.APIErr
	Update(context.Context, *users.User, string) apierrors.APIErr
	Delete(context.Context, string) (bool, apierrors.APIErr)
	FindByEmailAndPasswod(context.Context, string, string) (*users.User, apierrors.APIErr)
}

type userUseCase struct {
	userRepo    userrepo.UserRepository
	userService *userservice.UserService
}

//NewUserUseCase ...
func NewUserUseCase(userRepo userrepo.UserRepository, userService *userservice.UserService) UserUseCase {
	return &userUseCase{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (s *userUseCase) GetByEmail(ctx context.Context, email string) (*users.User, apierrors.APIErr) {

	var checkFormat bool

	email, checkFormat = s.userService.CheckEmail(email)
	if !checkFormat {
		return nil, apierrors.BadRequestAPIError("wrong email format")
	}

	user, getErr := s.userRepo.GetByEmail(ctx, email)

	if getErr != nil {
		return nil, getErr
	}

	return user, nil
}

func (s *userUseCase) Save(ctx context.Context, u *users.User) apierrors.APIErr {

	if u == nil {
		return apierrors.BadRequestAPIError("empty user")
	}

	if err := s.userService.ValidateUser(u); err != nil {
		return err
	}

	if err := s.userRepo.Save(ctx, u); err != nil {
		return err
	}

	return nil
}

func (s *userUseCase) FindByEmailAndPasswod(ctx context.Context, email string, password string) (*users.User, apierrors.APIErr) {

	var checkFormat bool

	email, checkFormat = s.userService.CheckEmail(email)

	if !checkFormat || password == "" {
		return nil, apierrors.BadRequestAPIError("email or password is empty")
	}

	password = s.userService.EncPassword(password)

	user, getErr := s.userRepo.FindByEmailAndPasswod(ctx, email, password)

	if getErr != nil {
		return nil, getErr
	}

	return user, nil
}

func (s *userUseCase) Delete(ctx context.Context, email string) (bool, apierrors.APIErr) {

	var checkFormat bool

	email, checkFormat = s.userService.CheckEmail(email)
	if !checkFormat {
		return false, apierrors.BadRequestAPIError("wrong email format")
	}

	confirm, err := s.userRepo.Delete(ctx, email)
	if err != nil {
		return false, err
	}

	return confirm, nil
}

func (s *userUseCase) Update(ctx context.Context, u *users.User, email string) apierrors.APIErr {

	if u == nil {
		return apierrors.BadRequestAPIError("empty user")
	}

	var checkFormat bool

	email, checkFormat = s.userService.CheckEmail(email)
	if !checkFormat {
		return apierrors.BadRequestAPIError("wrong email format")
	}

	currentUser, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	u = s.userService.UpdateUser(currentUser, u)

	if err := s.userService.ValidateUser(u); err != nil {
		return err
	}

	if err := s.userRepo.Update(ctx, u, email); err != nil {
		return err
	}

	return nil
}
