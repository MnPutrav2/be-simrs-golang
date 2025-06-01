package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserAccount struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,max=10"`
}

type UserTokenData struct {
	ID       int
	UserID   string
	CreateAt string
	Expired  string
}
