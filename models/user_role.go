package models

const RoleTableName = "role"

func (*UserRole) TableName() string {
	return RoleTableName
}

type UserRole struct {
	RoleID   uint   `gorm:"column:role_id;primaryKey" json:"role_id,omitempty"`
	RoleName string `gorm:"column:role_name;not null" json:"role_name,omitempty"`
	User    []User `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users,omitempty"`
}
