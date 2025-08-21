package dataLib

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name        string
	Email       string
	Age         uint8
	Birthday    time.Time
	MemberNum   sql.NullString
	activatedAt sql.NullTime
	CreatedAt   time.Time
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
	ignored     string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "-12123"
	return
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

	// 测试创建用户
	// user := &User{
	// 	Name: "测试用户",
	// 	Age:  25,
	// }

	// user.MemberNum.Valid = true
	// user.MemberNum.String = "M001"
	users := []*User{
		{Name: "小李", Age: 18, Birthday: time.Now()},
		{Name: "小张", Age: 19, Birthday: time.Now()},
	}

	result := db.Create(users)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("批量创建用户成功，创建了 %d 条记录\n", len(users))
	for i, user := range users {
		fmt.Printf("用户 %d: ID=%d, Name=%s, CreatedAt=%v\n",
			i+1, user.ID, user.Name, user.CreatedAt)
	}
}

func queryBlock() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())
	fmt.Println(block.Number().String())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty().String())
	fmt.Println(block.Hash().Hex())
	fmt.Println(block.Transactions())
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)

	// fmt.Println(block.Transactions()[0].Hash().Hex())
	// fmt.Println(block.Transactions()[0].To().Hex())
	// fmt.Println(block.Transactions()[0].From().Hex())

}
