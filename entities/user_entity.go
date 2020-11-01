package entities

type User struct {
	ID              int64  `json:"ID"`
	FirstName       string `json:"FIRST_NAME"`
	LastName        string `json:"LAST_NAME"`
	PhoneNum        string `json:"PHONE_NUM"`
	Email           string `json:"EMAIL"`
	Password        string `json:"PASSWORD"`
	ConfirmPassword string `json:"CONFIRM_PASSWORD"`
	CreateDate      string `json:"CREATE_DATE"`
	CreateBy        string `json:"CREATE_BY"`
}
