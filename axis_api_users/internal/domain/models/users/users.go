package users

//User ...
type User struct {
	ID          int64
	Email       string
	LastName    string
	FirstName   string
	Password    string
	Status      string
	JobTitle    string
	Privilege   string
	DoB         string
	PhotoURL    string
	DataCreated string
}

//NewUser ...
func NewUser(id int64, email string, lastName string, firstName string, password string,
	status string, jobTitle string, privilege string, dob string, photoURL string, dataCreated string) *User {

	return &User{
		ID:          id,
		Email:       email,
		LastName:    lastName,
		FirstName:   firstName,
		Password:    password,
		Status:      status,
		JobTitle:    jobTitle,
		Privilege:   privilege,
		DoB:         dob,
		PhotoURL:    photoURL,
		DataCreated: dataCreated,
	}
}
