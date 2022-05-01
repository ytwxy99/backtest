package database

import (
	"context"

	"gorm.io/gorm"
)

// add one history_day
func (historyDay *HistoryDay) AddHistoryDay(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(historyDay)
	return tx.Error
}

// get all history_day
func GetAllHistoryDay(ctx context.Context) ([]HistoryDay, error) {
	var historyDays []HistoryDay
	tx := ctx.Value("DbSession").(*gorm.DB).Find(&historyDays)
	return historyDays, tx.Error
}
