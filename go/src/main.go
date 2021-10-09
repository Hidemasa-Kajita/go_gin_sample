package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/Hidemasa-Kajita/go_api_sample/database"
    "github.com/Hidemasa-Kajita/go_api_sample/entity"
    "github.com/Hidemasa-Kajita/go_api_sample/repository"
)

func main() {
    db := database.Connect()

    // user 作成
    user1 := entity.User{Name: "山田太郎", Age: 25, Sex: "男"}
    user2 := entity.User{Name: "山田太郎", Age: 25, Sex: "男"}
    insertUsers := []entity.User{user1, user2}
    repository.Insert(insertUsers, db)

    // user 取得
    users := repository.FindAll(db)

    engine:= gin.Default()

    // user一覧を取得するAPI
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "users": users,
        })
    })

    db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&entity.User{})
    // defer db.Close() // 必要？

    engine.Run(":8080")
}
