package database

import (
	"fmt"
	"log"

	"github.com/Hidemasa-Kajita/go_api_sample/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
    // データベース
    Dialect = "mysql"

    // ユーザー名
    DBUser = "go"

    // パスワード
    DBPass = "go"

    // プロトコル
    DBProtocol = "tcp(mysql:3306)"

    // DB名
    DBName = "go"
)

func Connect() *gorm.DB {
    connectTemplate := "%s:%s@%s/%s?parseTime=true"
    connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
    db, err := gorm.Open(Dialect, connect)

    if err != nil {
        log.Println(err.Error())
    }

    db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&entity.User{})

    return db
}

func Disconnect(db *gorm.DB) {
    defer db.Close()
}