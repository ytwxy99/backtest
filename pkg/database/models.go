package database

import (
	"time"

	"gorm.io/gorm"
)

type Publish struct {
	gorm.Model

	Contract string
	Event    string
	Status   string
}

func (publish Publish) TableName() string {
	return "publish"
}

type HistoryDay struct {
	Contract string    `gorm:"primary_key;index:contract_idx;type:varchar(32)"`
	Time     time.Time `gorm:"primary_key;index:time_idx"`
	Price    string    `gorm:"type:varchar(32)"`
}

func (historyDay HistoryDay) TableName() string {
	return "history_day"
}

type HistoryFourHour struct {
	Contract string    `gorm:"primary_key;index:contract_idx;type:varchar(32)"`
	Time     time.Time `gorm:"primary_key;index:time_idx"`
	Price    string    `gorm:"type:varchar(32)"`
}

func (historyFourHour HistoryFourHour) TableName() string {
	return "history_four_hour"
}

type HistoryThirtyMin struct {
	Contract string    `gorm:"primary_key;index:contract_idx;type:varchar(32)"`
	Time     time.Time `gorm:"primary_key;index:time_idx"`
	Price    string    `gorm:"type:varchar(32)"`
}

func (historyThirtyMin HistoryThirtyMin) TableName() string {
	return "history_thirty_min"
}

type Order struct {
	gorm.Model

	Contract      string
	Fee_currency  string
	Price         string
	SoldPrice     string
	Amount        float32
	Time          int64
	Tp            float32
	Sl            float32
	Ttp           float32
	Tsl           float32
	Text          string
	Status        string
	Typee         string
	Account       string
	Side          string
	Iceberg       string
	Left          float32
	Fee           float32
	Amount_filled float32
	Direction     string
	BuyTime       time.Time
	SoldTime      time.Time
}

func (order Order) TableName() string {
	return "orders"
}
