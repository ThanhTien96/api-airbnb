package models

const UserTableName = "user"

func (*UserBase) TableName() string {
	return UserTableName
}

type UserBase struct {
	BaseModel         `json:",omitempty"`
	CreateUserRequest `json:",omitempty"`
	UserID            uint `gorm:"column:user_id;primaryKey" json:"user_id,omitempty"`
}

type CreateUserRequest struct {
	UserName    string  `gorm:"column:username;index;not null" json:"username,omitempty"`
	Password    string  `gorm:"column:password;not null" json:"password,omitempty"`
	Email       string  `gorm:"column:email;not null" json:"email,omitempty"`
	PhoneNumber string  `gorm:"column:phone_number;not null" json:"phone_number,omitempty"`
	Avatar      string `gorm:"column:avatar" json:"avatar,omitempty"`
	RoleID      uint	`gorm:"column:role_id" json:"role_id,omitempty"`
}

type User struct {
	UserBase UserBase `gorm:"embedded" json:",omitempty"`
	Role *UserRole `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
