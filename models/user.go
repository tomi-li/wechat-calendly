package models

import (
	"calendly/db"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID(id string) (*User, error) {
	user := User{ID: id}
	db.GetMysql().Create(&user)
	return &user, nil
}

func (h User) ListAll() []User {
	var users []User
	db.GetMysql().Find(&users)
	return users
}
