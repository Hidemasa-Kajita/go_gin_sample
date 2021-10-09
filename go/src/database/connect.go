package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
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
    connectTemplate := "%s:%s@%s/%s"
    connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
    db, err := gorm.Open(Dialect, connect)

    if err != nil {
        log.Println(err.Error())
    }

    return db
}