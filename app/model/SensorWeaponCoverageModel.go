package model

type SensorWeaponCoverage struct {
	ID           int64   `gorm:"primaryKey;autoIncrement"`
	Type         string  `gorm:"type:varchar(255)"`
	GroupSWC     string  `gorm:"column:group_swc;type:varchar(255)"`
	Item         string  `gorm:"type:varchar(255)"`
	Environment  string  `gorm:"type:varchar(255)"`
	Value        float64 `gorm:"type:double precision"`
	DefaultValue float64 `gorm:"column:default_value;type:double precision"`
	Unit         string  `gorm:"type:varchar(255)"`
	IsDeleted    bool    `gorm:"column:is_deleted"`
	CreatedAt    string  `gorm:"column:created_at"`
	UpdatedAt    string  `gorm:"column:updated_at"`
}

type SWCParameterResult struct {
	GroupSWC     string `json:"group_swc"`
	Environment  string `json:"environment"`
	Type         string `json:"type"`
	ID           string `json:"id"`
	Value        string `json:"value"`
	DefaultValue string `json:"default_value"`
	Item         string `json:"item"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	IsDeleted    string `json:"is_deleted"`
	Unit         string `json:"unit"`
}

func (SensorWeaponCoverage) TableName() string {
	return "swc_parameter"
}
