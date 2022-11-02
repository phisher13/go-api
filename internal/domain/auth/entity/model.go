package entity

type UserModel struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
