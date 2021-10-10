package repository

import (
	"github.com/Hidemasa-Kajita/go_api_sample/database"
	"github.com/Hidemasa-Kajita/go_api_sample/entity"
)

func Insert(user entity.User) entity.User {
    db := database.Connect()
    db.NewRecord(user)
    db.Create(&user)

    return user
}

func FindAll() []entity.User {
    db := database.Connect()
    defer db.Close()
    var users []entity.User
    db.Find(&users)

    return users
}

func FindOneById(id string) entity.User {
    db := database.Connect()
    var user entity.User
    db.First(&user, id)
    return user
}

// idとかが入った`user`構造体を取得できない。
func Update(id string, user entity.User) entity.User {
    db := database.Connect()
    db.Model(&user).Where("id = ?", id).Update(user)
    return user
}

func Delete(id string) {
    db := database.Connect()
    db.Delete(&entity.User{}, id)
}
