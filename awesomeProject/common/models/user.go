package models

type User struct {
	ID       int    `json:"id" db:"id" gorm:"primary_key"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
