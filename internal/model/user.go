package model

type User struct {
	CustomBaseModel
	Username  string   `gorm:"not null" json:"username"`
	Email     string   `gorm:"type:citext;unique;not null" json:"email"`
	Languages []string `gorm:"type:varchar(100)" json:"languages"`
}
