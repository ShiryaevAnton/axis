package store

import (
	"database/sql"
	"errors"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apistore"
	"github.com/ShiryaevAnton/axis1/axis_api_utils/apierrors"
)

var (
	store Store
)

type storeImp struct {
	db *sql.DB
}

func newStore(db *sql.DB) Store {
	return &storeImp{db: db}
}

//NewStore ...
func NewStore(username string, password string, host string, port string, dbname string) error {
	db, err := apistore.InitDatabase(username, password, host, port, dbname)
	if err != nil {
		return err
	}

	store = newStore(db)

	return nil
}

//GetStore ...
func GetStore() Store {
	return store
}

func (s *storeImp) Get(param string, query string) (*users.User, apierrors.APIErr) {

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, apierrors.InternalServerAPIError("error when trying to prepare to get user", errors.New("database error"))
	}

	defer stmt.Close()

	var user users.User

	if getErr := stmt.QueryRow(param).Scan(
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
		return nil, getErr
	}

	return &user, nil
}

func (s *storeImp) Save(u *users.User) apierrors.APIErr {
	return nil
}

func (s *storeImp) Update(u *users.User) apierrors.APIErr {
	return nil
}

func (s *storeImp) Delete(string) (bool, apierrors.APIErr) {
	return false, nil
}
