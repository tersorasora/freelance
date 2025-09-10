package entity

type User struct {
	UserID 		string `gorm:"column:user_id;primaryKey"`
    Email  		string `gorm:"column:email;unique;not null"`
	Name   		string `gorm:"column:name;not null"`
	Password 	string `gorm:"column:password;not null"`
	Balance 	float64 `gorm:"column:balance;not null"`
	RoleID 		string `gorm:"column:role_id;not null"`
}