package entity

type Contract struct {
	ContractID 			string `gorm:"column:contract_id;primaryKey"`
    ContractStart  		string `gorm:"column:contract_start;not null"`
	ContractEnd 		string `gorm:"column:contract_end;not null"`
	ServiceID 			string `gorm:"column:service_id;not null"`
	UserID 				string `gorm:"column:user_id;not null"`
}