package repository

import (
    "github.com/jinzhu/gorm"
	"github.com/Hidemasa-Kajita/go_api_sample/entity"
)

func Insert(users []entity.User, db *gorm.DB) {
    for _, user := range users {
        db.NewRecord(user)
        db.Create(&user)
    }
}

func FindAll(db *gorm.DB) []entity.User {
    var allUsers []entity.User
    db.Find(&allUsers)
    return allUsers
}
