package entity

type User struct {
	UserID 		string `gorm:"column:user_id;primaryKey"`
    Email  		string `gorm:"column:email;unique;not null"`
	Name   		string `gorm:"column:name;not null"`
	Password 	string `gorm:"column:password;not null"`
}