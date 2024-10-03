package types

type User struct {
	ID           string  `json:"id" gorm:"primaryKey;autoIncrement"`
	GoogleID     *string `gorm:"unique"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    int64   `json:"expires_in"`
}
