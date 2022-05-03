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
