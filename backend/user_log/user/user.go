// type User struct {
// 	models.BaseModel

// 	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
// 	Email           string `gorm:"type:varchar(255);unique;" valid:"email"`
// 	Password        string `gorm:"type:varchar(255)" valid:"password"`
// 	ActivationToken string `gorm:"varchar(255)"`
// 	Activated       int    `gorm:"type:int(10)"`
// 	EmailVerifiedAt *time.Time

// 	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
// 	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
// }
