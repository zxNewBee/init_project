package dataLib

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Account struct {
	ID      uint    `gorm:"primarykey"`
	Balance float64 `gorm:"not null;default:0"`
}

type Transactions struct {
	ID            uint    `gorm:"primarykey"`
	FromAccountId uint    `gorm:"not null;index"`
	ToAccountId   uint    `gorm:"not null;index"`
	Amount        float64 `gorm:"not null"`
}

// AutoMigrateAccountAndTransactions 自动创建/更新账户与交易记录表结构
func AutoMigrateAccountAndTransactions(db *gorm.DB) error {
	return db.AutoMigrate(&Account{}, &Transactions{})
}

// TransferFunds 在单个数据库事务中完成转账：余额检查、扣款、加款与交易记录
func TransferFunds(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("转账金额必须大于0")
	}
	if fromAccountID == toAccountID {
		return errors.New("转出与转入账户不能相同")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// 行级锁读取两条账户记录，防止并发修改
		var accounts []Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id IN ?", []uint{fromAccountID, toAccountID}).
			Find(&accounts).Error; err != nil {
			return err
		}
		if len(accounts) != 2 {
			return errors.New("账户不存在或数量不足")
		}

		var fromAccount, toAccount *Account
		for i := range accounts {
			if accounts[i].ID == fromAccountID {
				fromAccount = &accounts[i]
			} else if accounts[i].ID == toAccountID {
				toAccount = &accounts[i]
			}
		}
		if fromAccount == nil || toAccount == nil {
			return errors.New("无法定位转出或转入账户")
		}

		if fromAccount.Balance < amount {
			return errors.New("余额不足，转账失败")
		}

		// 更新余额（已持有 FOR UPDATE 锁）
		if err := tx.Model(&Account{}).
			Where("id = ?", fromAccountID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Model(&Account{}).
			Where("id = ?", toAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 记录交易流水
		tr := &Transactions{
			FromAccountId: fromAccountID,
			ToAccountId:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(tr).Error; err != nil {
			return err
		}

		return nil
	})
}
