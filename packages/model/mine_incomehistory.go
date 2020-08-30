package model

import (
	"github.com/shopspring/decimal"
	"time"
)

//MineIncomehistory example
type MineIncomehistory struct {
	ID                      int64           `gorm:"primary_key;not null"`
	Devid                   int64           `gorm:"not null"`
	Number                  string          `gorm:"not null" `
	Poolid                  int64           `gorm:"not null" `
	Keyid                   int64           `gorm:"not null"`
	Mineid                  int64           `gorm:"not null"`
	Amount                  decimal.Decimal `gorm:"not null"`
	Type                    int64           `gorm:"not null"`
	Capacitys               int64           `gorm:"not null" `
	Nonce                   int64           `gorm:"not null" `
	Foundation              decimal.Decimal `gorm:"not null" `
	Mine_incomehistory_hash []byte          `gorm:"not null`
	Block_id                int64           `gorm:"not null"`
	Date_created            int64           `gorm:"not null default 0"`
}

//DayMineIncomehistory example
type DayMineIncomehistory struct {
	Amount decimal.Decimal `json:"amount"`
	Time   int64           `json:"time"`
}

// TableName returns name of table
func (m MineIncomehistory) TableName() string {
	return `1_mine_incomehistory`
}

// Get is retrieving model from database
func (m *MineIncomehistory) Get(id int64) (bool, error) {
	return isFound(DBConn.Where("block_id = ?", id).First(m))
}

// Get is retrieving model from database
func (m *MineIncomehistory) GetDelay(id int64) (bool, error) {
	for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
		}
	}
	return isFound(DBConn.Where("block_id = ?", id).First(m))
}

// Get is retrieving model from database
func (m *MineIncomehistory) GetActiveMiner(time, devid int64) (incomes []MineIncomehistory, err error) {
	db := GetDB(nil)
	err = db.Table("1_mine_incomehistory").
		Where("devid = ? and date_created >?", devid, time).
		Order("date_created asc").
		Scan(&incomes).Error
	return incomes, err
}