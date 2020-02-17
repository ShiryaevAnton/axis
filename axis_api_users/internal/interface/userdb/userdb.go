package userdb

import (
	"context"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	userrepo "github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/repositories/userrepositories"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/store"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apierrors"
)

const (
	queryInserstUser            = "INSERT INTO users (first_name, last_name, email, job_title, privilege, dob, photo_url, status, password, data_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
	queryGetUserByEmail         = "SELECT first_name, last_name, email, job_title, privilege, dob, photo_url, status, data_created  FROM users WHERE email=$1;"
	queryUpdateUser             = "UPDATE users SET first_name=$1, last_name=$2, email=$3, job_title=$4, privilege=$5, dob=$6, photo_url=$7, status=$8, password=$9, data_created=$10 WHERE email=$11;"
	queryDeleteUser             = "DELETE FROM users WHERE email=$1;"
	queryFindUserByStatus       = "SELECT first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailandPassword = "SELECT first_name, last_name, email, job_title, privilege, dob, photo_url, status, data_created  FROM users WHERE email=$1 AND password=$2;"
)

//UserRepositoryImp ...
type userRepositoryImp struct {
	conn store.Store
}

//NewUserRepository ...
func NewUserRepository(conn store.Store) userrepo.UserRepository {
	return &userRepositoryImp{conn: conn}
}

//GetByID ...
func (r *userRepositoryImp) GetByEmail(ctx context.Context, email string) (*users.User, apierrors.APIErr) {

	//var user *users.User

	user, err := r.conn.Get(queryGetUserByEmail, email)
	if err != nil {

	}

	return nil, nil
}

func (r *userRepositoryImp) Save(ctx context.Context, u *users.User) apierrors.APIErr {

	return nil
}

func (r *userRepositoryImp) Update(ctx context.Context, u *users.User, email string) apierrors.APIErr {

	return nil
}

func (r *userRepositoryImp) Delete(ctx context.Context, email string) (bool, apierrors.APIErr) {
	return false, nil

}

//FindByEmailAndPasswod ...
func (r *userRepositoryImp) FindByEmailAndPasswod(ctx context.Context, email string, password string) (*users.User, apierrors.APIErr) {

	return nil, nil
}
