package userservice

import (
	"strings"

	"github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/models/users"
	"github.com/ShiryaevAnton/axis/axis_api_users/utils"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apierrors"
)

const (
	userPrivilege     = "user"
	adminPrivilege    = "admin"
	managerPrivilege  = "manager"
	notApprovedStatus = "not_approved"
	activeStatus      = "active"
	deactiveStatus    = "deactive"
	techJobTitle      = "technician"
	managerJobTitle   = "manager"
	defaultPhotoURL   = "defaultPhotoURL" //TODO
)

//UserService ...
type UserService struct{}

//NewUserService ...
func NewUserService() *UserService {
	return &UserService{}
}

//ValidateUser ...
func (s *UserService) ValidateUser(u *users.User) apierrors.APIErr {

	u.LastName = strings.TrimSpace(u.LastName)
	if u.LastName == "" {
		return apierrors.BadRequestAPIError("invalid last name")
	}

	u.FirstName = strings.TrimSpace(u.FirstName)
	if u.FirstName == "" {
		return apierrors.BadRequestAPIError("invalid first name")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return apierrors.BadRequestAPIError("invalid password")
	}
	u.Password = utils.GetSHA256(u.Password)

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if !utils.CheckEmail(u.Email) {
		return apierrors.BadRequestAPIError("invalid email")
	}

	u.DataCreated = utils.GetNowString()

	u.Privilege = strings.TrimSpace(u.Privilege)
	if u.Privilege == "" && u.Privilege != adminPrivilege && u.Privilege != managerPrivilege {
		u.Privilege = userPrivilege
	}

	u.Status = strings.TrimSpace(u.Status)
	if u.Status == "" && u.Status != activeStatus && u.Status != deactiveStatus {
		u.Status = notApprovedStatus
	}

	u.JobTitle = strings.TrimSpace(u.JobTitle)
	if u.JobTitle != techJobTitle && u.JobTitle != managerJobTitle {
		u.JobTitle = ""
	}

	u.PhotoURL = strings.TrimSpace(u.PhotoURL)
	if !utils.CheckURL(u.PhotoURL) {
		u.PhotoURL = defaultPhotoURL
	}

	u.DoB = strings.TrimSpace(u.DoB)
	if !utils.CheckURL(u.DoB) {
		u.DoB = ""
	}

	return nil
}

//CheckEmail ...
func (s *UserService) CheckEmail(email string) (string, bool) {

	email = strings.TrimSpace(strings.ToLower(email))

	return email, utils.CheckEmail(email)
}

//EncPassword ...
func (s *UserService) EncPassword(password string) string {
	return utils.GetSHA256(password)
}

//UpdateUser ...
func (s *UserService) UpdateUser(currentUser *users.User, u *users.User) *users.User {

	if u.FirstName != "" {
		currentUser.FirstName = u.FirstName
	}

	if u.LastName != "" {
		currentUser.LastName = u.LastName
	}

	if u.Email != "" {
		currentUser.Email = u.Email
	}

	if u.DoB != "" {
		currentUser.DoB = u.DoB
	}

	if u.JobTitle != "" {
		currentUser.JobTitle = u.JobTitle
	}

	if u.PhotoURL != "" {
		currentUser.PhotoURL = u.PhotoURL
	}

	if u.Privilege != "" {
		currentUser.Privilege = u.Privilege
	}

	if u.Status != "" {
		currentUser.Status = u.Status
	}

	if u.Password != "" {
		currentUser.Password = s.EncPassword(u.Password)
	}

	return currentUser
}
