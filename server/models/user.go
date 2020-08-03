package models

import (
	"github.com/jinzhu/gorm"
)

// User ...user struct
type User struct {
	gorm.Model
	Name string `gorm:"type: varchar(100)"`
}
