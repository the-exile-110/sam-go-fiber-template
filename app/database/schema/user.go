package schema

import "time"

type Users struct {
	Id            int       `json:"id" gorm:"primary_key,type:int,omitempty,auto_increment:true"`
	UserId        string    `json:"userId" gorm:"type:varchar(255);not null"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
	Email         string    `json:"email" gorm:"type:varchar(255);not null"`
	Birthdate     string    `json:"birthdate" gorm:"type:varchar(255);not null"`
	Locale        string    `json:"locale" gorm:"type:varchar(255);not null"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"type:varchar(255);not null"`
	LastLoginDate time.Time `json:"lastLoginDate" gorm:"type:datetime;not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
