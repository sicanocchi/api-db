package model

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (user User) TableName() string {
	return "users"
}

func (u *User) PrepareGive() {
	u.Password = ""
}
