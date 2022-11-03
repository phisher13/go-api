package entity

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ProductModel struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
