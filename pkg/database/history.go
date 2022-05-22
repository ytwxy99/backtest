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

// AddHistoryFourHour add a 4h history data
func (historyFourHour *HistoryFourHour) AddHistoryFourHour(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(historyFourHour)
	return tx.Error
}

// FetchHistoryFourHour get specified coin 4h history prices
func (historyFourHour *HistoryFourHour) FetchHistoryFourHour(ctx context.Context) ([]*HistoryFourHour, error) {
	histoies := []*HistoryFourHour{}
	rows, err := ctx.Value("DbSession").(*gorm.DB).Table("history_four_hour").
		Where("contract = ?", historyFourHour.Contract).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		history := &HistoryFourHour{}
		if err := ctx.Value("DbSession").(*gorm.DB).ScanRows(rows, history); err != nil {
			return nil, err
		}

		histoies = append(histoies, history)
	}

	return histoies, nil
}

// AddHistoryThirtyMin add a 30m history data
func (historyThirtyMin *HistoryThirtyMin) AddHistoryThirtyMin(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(historyThirtyMin)
	return tx.Error
}

// FetchHistoryThirtyMin get specified coin 30m history prices
func (historyThirtyMin *HistoryThirtyMin) FetchHistoryThirtyMin(ctx context.Context) ([]*HistoryThirtyMin, error) {
	histoies := []*HistoryThirtyMin{}
	rows, err := ctx.Value("DbSession").(*gorm.DB).Table("history_thirty_min").
		Where("contract = ?", historyThirtyMin.Contract).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		history := &HistoryThirtyMin{}
		if err := ctx.Value("DbSession").(*gorm.DB).ScanRows(rows, history); err != nil {
			return nil, err
		}

		histoies = append(histoies, history)
	}

	return histoies, nil
}
