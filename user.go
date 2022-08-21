package authy

type User struct {
	ID          string `json:"id"`
	AuthyID     string `json:"authy_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}
