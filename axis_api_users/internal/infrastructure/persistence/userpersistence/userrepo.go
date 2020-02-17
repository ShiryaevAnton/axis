package userpersistence

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	userrepo "github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/repositories/userrepositories"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apierrors"
)

const (
	queryInserstUser            = "INSERT INTO users (first_name, last_name, email, job_title, privilege, dob, photo_url, status, password, data_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
	queryGetUser                = "SELECT first_name, last_name, email, job_title, privilege, dob, photo_url, status, data_created  FROM users WHERE email=$1;"
	queryUpdateUser             = "UPDATE users SET first_name=$1, last_name=$2, email=$3, job_title=$4, privilege=$5, dob=$6, photo_url=$7, status=$8, password=$9, data_created=$10 WHERE email=$11;"
	queryDeleteUser             = "DELETE FROM users WHERE email=$1;"
	queryFindUserByStatus       = "SELECT first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailandPassword = "SELECT first_name, last_name, email, job_title, privilege, dob, photo_url, status, data_created  FROM users WHERE email=$1 AND password=$2;"
)

//UserRepositoryImp ...
type userRepositoryImp struct {
	conn *sql.DB
}

//NewUserRepository ...
func NewUserRepository(conn *sql.DB) userrepo.UserRepository {
	return &userRepositoryImp{conn: conn}
}

//GetByID ...
func (r *userRepositoryImp) GetByEmail(ctx context.Context, email string) (*users.User, apierrors.APIErr) {

	stmt, err := r.conn.Prepare(queryGetUser)
	if err != nil {
		return nil, apierrors.InternalServerAPIError("error when trying to prepare to get user", errors.New("database error"))
	}

	defer stmt.Close()

	var user users.User

	if getErr := stmt.QueryRow(email).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.JobTitle,
		&user.Privilege,
		&user.DoB,
		&user.PhotoURL,
		&user.Status,
		&user.DataCreated,
	); getErr != nil {
		//TODO if not found
		return nil, apierrors.InternalServerAPIError("error when trying to execute getting user", errors.New("database error"))
	}

	return &user, nil
}

func (r *userRepositoryImp) Save(ctx context.Context, u *users.User) apierrors.APIErr {

	stmt, err := r.conn.Prepare(queryInserstUser)
	if err != nil {
		return apierrors.InternalServerAPIError("error when trying to prepate to save user", errors.New("database error"))
	}

	defer stmt.Close()

	_, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.JobTitle, u.Privilege,
		u.DoB, u.PhotoURL, u.Status, u.Password, u.DataCreated)

	if saveErr != nil {
		return apierrors.InternalServerAPIError("error when trying to execute saving user", errors.New("database error"))
	}

	return nil
}

func (r *userRepositoryImp) Update(ctx context.Context, u *users.User, email string) apierrors.APIErr {

	stmt, err := r.conn.Prepare(queryUpdateUser)
	if err != nil {
		return apierrors.InternalServerAPIError("error when trying to prepate to update user", errors.New("database error"))
	}

	defer stmt.Close()

	//Fix me
	_, updErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.JobTitle, u.Privilege,
		u.DoB, u.PhotoURL, u.Status, u.Password, u.DataCreated, email)

	if updErr != nil {
		return apierrors.InternalServerAPIError("error when trying to execute updating user", errors.New("database error"))
	}

	return nil
}

func (r *userRepositoryImp) Delete(ctx context.Context, email string) (bool, apierrors.APIErr) {

	stmt, err := r.conn.Prepare(queryDeleteUser)
	if err != nil {
		return false, apierrors.InternalServerAPIError("error when trying to prepate to delete user", errors.New("database error"))
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(email)
	if delErr != nil {
		return false, apierrors.InternalServerAPIError("error when trying to execute deleting user", errors.New("database error"))
	}

	return true, nil
}

//FindByEmailAndPasswod ...
func (r *userRepositoryImp) FindByEmailAndPasswod(ctx context.Context, email string, password string) (*users.User, apierrors.APIErr) {

	stmt, err := r.conn.Prepare(queryFindByEmailandPassword)
	if err != nil {
		return nil, apierrors.InternalServerAPIError("error when trying to prepare to find user by email and password", errors.New("database error"))
	}

	defer stmt.Close()

	var user users.User

	if getErr := stmt.QueryRow(email, password).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.JobTitle,
		&user.Privilege,
		&user.DoB,
		&user.PhotoURL,
		&user.Status,
		&user.DataCreated,
	); getErr != nil {
		//TODO if not found
		return nil, apierrors.InternalServerAPIError("error when trying to get user", errors.New("database error"))
	}

	return &user, nil
}
