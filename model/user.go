package model

type User struct {
	Id       string    `gorm:"primaryKey;type:varchar(255)"`
	Username string    `gorm:"not null;type:varchar(255)"`
	Password string    `gorm:"not null;type:varchar(255)"`
	Product  []Product `gorm:"foreignKey:UserId"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Id string `json:"id"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
