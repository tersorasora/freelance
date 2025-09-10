package entity

type Field struct {
	FieldID   string `gorm:"column:field_id;primaryKey"`
	FieldName string `gorm:"column:field_name;unique;not null"`
}