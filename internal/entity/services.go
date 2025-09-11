package entity

type Service struct {
	ServiceID   string    `gorm:"column:service_id;primaryKey"`
	ServiceName string    `gorm:"column:service_name;not null"`
	Description string    `gorm:"column:description;not null"`
	Price       float64   `gorm:"column:price;not null"`
	Period      string	  `gorm:"column:period;not null"`
	FieldID     string    `gorm:"column:field_id;not null"`
	UserID      string    `gorm:"column:user_id;not null"`
}