package dataLib

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name        string
	Email       *string
	Age         uint8
	Birthday    *time.Time
	MemberNum   sql.NullString
	activatedAt sql.NullTime
	CreatedAt   time.Time
	UpdateAt    time.Time
	ignored     string
}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}
type Blog2 struct {
	ID      int64
	Upvotes int32
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Blog2{})

}
