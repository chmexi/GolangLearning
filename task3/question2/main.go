package main

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Accounts struct {
	ID      uint
	Balance float32
}

type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float32
}

func APaysToB(db *gorm.DB, a Accounts, b Accounts, amount float32) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&a, a.ID).Error; err != nil {
			return err
		}
		if err := tx.First(&b, b.ID).Error; err != nil {
			return err
		}
		if a.Balance < amount {
			return fmt.Errorf("余额不足")
		}

		if err := tx.Debug().Model(&a).Update("balance", a.Balance-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&b).Update("balance", b.Balance+amount).Error; err != nil {
			return err
		}
		transaction := Transaction{
			FromAccountID: a.ID,
			ToAccountID:   b.ID,
			Amount:        amount,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err

		}
		return nil
	})
	return err
}

func main() {
	dsn := "root:meng0987612345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error cannot open database error: ", err)
	}

	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transaction{})

	a := Accounts{ID: 1, Balance: 10}
	b := Accounts{ID: 2, Balance: 9}
	// db.Create([]Accounts{a, b})
	accounts := []Accounts{}
	db.Find(&accounts)
	fmt.Println("转账前", accounts)
	err = APaysToB(db, a, b, 100)
	if err != nil {
		fmt.Println("转账失败 err:", err)
	}
	accounts = []Accounts{}
	db.Find(&accounts)
	fmt.Println("转账后", accounts)
}
