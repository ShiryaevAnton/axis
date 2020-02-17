package userserver

import (
	"context"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	usergrpcpb "github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/usergrpc/proto"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/usecases/userusecases"
)

//UserServer ...
type UserServer struct {
	userUseCase userusecases.UserUseCase
}

//NewUserServer ...
func NewUserServer(userUseCase userusecases.UserUseCase) *UserServer {
	return &UserServer{
		userUseCase: userUseCase,
	}
}

//GetUser ...
func (s *UserServer) GetUser(ctx context.Context, req *usergrpcpb.GetUserRequest) (*usergrpcpb.GetUserResponse, error) {

	email := req.GetEmail()

	user, err := s.userUseCase.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	response := &usergrpcpb.GetUserResponse{
		User: &usergrpcpb.User{
			Email:       user.Email,
			LastName:    user.LastName,
			FirstName:   user.FirstName,
			Status:      user.Status,
			JobTitle:    user.JobTitle,
			Privilege:   user.Privilege,
			PhotoUrl:    user.PhotoURL,
			Dob:         user.DoB,
			DataCreated: user.DataCreated,
		},
	}

	return response, nil
}

//SaveUser ...
func (s *UserServer) SaveUser(ctx context.Context, req *usergrpcpb.SaveUserRequest) (*usergrpcpb.SaveUserResponse, error) {

	user := users.User{
		Email:     req.GetUserSave().GetEmail(),
		LastName:  req.GetUserSave().GetLastName(),
		FirstName: req.GetUserSave().GetFirstName(),
		Status:    req.GetUserSave().GetStatus(),
		JobTitle:  req.GetUserSave().GetJobTitle(),
		Privilege: req.GetUserSave().GetPrivilege(),
		PhotoURL:  req.GetUserSave().GetPhotoUrl(),
		Password:  req.GetUserSave().GetPassword(),
	}

	if err := s.userUseCase.Save(ctx, &user); err != nil {
		return nil, err
	}

	response := &usergrpcpb.SaveUserResponse{
		User: &usergrpcpb.User{
			Email:       user.Email,
			LastName:    user.LastName,
			FirstName:   user.FirstName,
			Status:      user.Status,
			JobTitle:    user.JobTitle,
			Privilege:   user.Privilege,
			PhotoUrl:    user.PhotoURL,
			Dob:         user.DoB,
			DataCreated: user.DataCreated,
		},
	}

	return response, nil
}

//DeleteUser ...
func (s *UserServer) DeleteUser(ctx context.Context, req *usergrpcpb.DeleteUserRequest) (*usergrpcpb.DeleteUserResponse, error) {

	email := req.GetEmail()

	confirm, err := s.userUseCase.Delete(ctx, email)
	if err != nil {
		return nil, err
	}

	return &usergrpcpb.DeleteUserResponse{
		ConfirmDelete: confirm,
	}, nil
}

//UpdateUser ...
func (s *UserServer) UpdateUser(ctx context.Context, req *usergrpcpb.UpdateUserRequest) (*usergrpcpb.UpdateUserResponse, error) {

	email := req.GetEmail()

	user := users.User{
		Email:     req.GetUserSave().GetEmail(),
		LastName:  req.GetUserSave().GetLastName(),
		FirstName: req.GetUserSave().GetFirstName(),
		Status:    req.GetUserSave().GetStatus(),
		JobTitle:  req.GetUserSave().GetJobTitle(),
		Privilege: req.GetUserSave().GetPrivilege(),
		PhotoURL:  req.GetUserSave().GetPhotoUrl(),
		Password:  req.GetUserSave().GetPassword(),
	}

	if err := s.userUseCase.Update(ctx, &user, email); err != nil {
		return nil, err
	}

	response := &usergrpcpb.UpdateUserResponse{
		User: &usergrpcpb.User{
			Email:       user.Email,
			LastName:    user.LastName,
			FirstName:   user.FirstName,
			Status:      user.Status,
			JobTitle:    user.JobTitle,
			Privilege:   user.Privilege,
			PhotoUrl:    user.PhotoURL,
			Dob:         user.DoB,
			DataCreated: user.DataCreated,
		},
	}

	return response, nil
}

//FindByEmailAndPassword ...
func (s *UserServer) FindByEmailAndPassword(ctx context.Context, req *usergrpcpb.FindByEmailAndPasswordRequest) (*usergrpcpb.FindByEmailAndPasswordResponse, error) {

	email := req.GetEmail()
	password := req.GetPassword()

	user, err := s.userUseCase.FindByEmailAndPasswod(ctx, email, password)
	if err != nil {
		return nil, err
	}

	response := &usergrpcpb.FindByEmailAndPasswordResponse{
		User: &usergrpcpb.User{
			Email:       user.Email,
			LastName:    user.LastName,
			FirstName:   user.FirstName,
			Status:      user.Status,
			JobTitle:    user.JobTitle,
			Privilege:   user.Privilege,
			PhotoUrl:    user.PhotoURL,
			Dob:         user.DoB,
			DataCreated: user.DataCreated,
		},
	}

	return response, nil
}
