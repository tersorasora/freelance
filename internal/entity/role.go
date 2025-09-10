package entity

type Role struct {
	RoleID 			string `gorm:"column:role_id;primaryKey"`
    RoleName  		string `gorm:"column:role_name;unique;not null"`
}