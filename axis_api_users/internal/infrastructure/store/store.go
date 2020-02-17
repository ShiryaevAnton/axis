package store

import (
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	"github.com/ShiryaevAnton/axis1/axis_api_utils/apierrors"
)

//Store ...
type Store interface {
	Get(string, string) (*users.User, apierrors.APIErr)
	Save(*users.User) apierrors.APIErr
	Delete(string) (bool, apierrors.APIErr)
	Update(*users.User) apierrors.APIErr
	//GetAll ()
}
