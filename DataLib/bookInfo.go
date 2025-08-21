package dataLib

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type BookInfo struct {
	BookId   int    `gorm:"primarykey"` // 图书ID
	BookName string `json:"book_name"`  // 图书名称
	Author   string `json:"author"`     // 作者
	Price    int    `json:"price"`      // 价格
	Title    string `json:"title"`      // 书名
}

func CreateBookTable(db *sqlx.DB) BookInfo {
	var book BookInfo
	if db == nil {
		fmt.Println("gormDb is nil")
		return book

	} else {
		query := `CREATE TABLE IF NOT EXISTS book_infos (
			BookId INT PRIMARY KEY AUTO_INCREMENT,
			BookName VARCHAR(100) NOT NULL,
			Author VARCHAR(100) NOT NULL,
			Price INT NOT NULL,
			Title VARCHAR(100) NOT NULL
		
		)`
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.QueryRowx(query)
		queryInsert := `INSERT INTO book_Infos (BookName,Author,Price,Title) 
		VALUES("ABC","ZX",190,"AAA")
		`
		db.ExecContext(ctx, queryInsert)

		// if _, err := db.ExecContext(ctx, query); err != nil {
		// 	fmt.Println("create table book_infos failed:", err)
		// 	return book
		// }

	}

	return book
}
func queryBookTable(db *gorm.DB) {

}
