package model

type Product struct {
	Id          string `gorm:"primaryKey;type:varchar(255)"`
	ProductName string `gorm:"not null;type:varchar(255)"`
	Price       int    `gorm:"not null"`
	UserId      string
}

type ProductCreateRequest struct {
	Price int `json:"price"`
}

type ProductCreateResponse struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Price  int    `json:"price"`
}
