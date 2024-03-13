package models


type BaseModel struct {
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at,omitempty"`
}




