package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
	Name string `gorm:"size:255"`
    Age  int
    Sex  string `gorm:"size:255"`
}

func (u User) String() string {
    return fmt.Sprintf("%s(%d)", u.Name, u.Age)
}
