package database

import (
	"context"

	"gorm.io/gorm"
)

// AddHistoryDay add one history_day
func (historyDay *HistoryDay) AddHistoryDay(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(historyDay)
	return tx.Error
}

// GetAllHistoryDay get all history_day
func GetAllHistoryDay(ctx context.Context) ([]HistoryDay, error) {
	var historyDays []HistoryDay
	tx := ctx.Value("DbSession").(*gorm.DB).Find(&historyDays)
	return historyDays, tx.Error
}

// FetchHistoryDay get specified coin history prices
func (historyDay *HistoryDay) FetchHistoryDay(ctx context.Context) ([]*HistoryDay, error) {
	histoies := make([]*HistoryDay, 0)
	rows, err := ctx.Value("DbSession").(*gorm.DB).Table("history_day").
		Where("contract = ?", historyDay.Contract).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		history := &HistoryDay{}
		if err := ctx.Value("DbSession").(*gorm.DB).ScanRows(rows, history); err != nil {
			return nil, err
		}

		histoies = append(histoies, history)
	}

	return histoies, nil
}