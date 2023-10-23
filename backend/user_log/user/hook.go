// package user

// import (

//     "goblog/pkg/strings"
//     "gorm.io/gorm"
// )

// // BeforeCreate GORM 的模型钩子，创建模型前调用
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

//     u.ActivationToken = strings.Rand(10)
//     return
// }