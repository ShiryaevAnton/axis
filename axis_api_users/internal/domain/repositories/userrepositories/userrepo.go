package userrepo

import (
	"context"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apierrors"
)

//UserRepository ...
type UserRepository interface {
	GetByEmail(context.Context, string) (*users.User, apierrors.APIErr)
	Save(context.Context, *users.User) apierrors.APIErr
	Update(context.Context, *users.User, string) apierrors.APIErr
	Delete(context.Context, string) (bool, apierrors.APIErr)
	FindByEmailAndPasswod(context.Context, string, string) (*users.User, apierrors.APIErr)
}
