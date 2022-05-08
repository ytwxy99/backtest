package database

import (
	"context"
	"gorm.io/gorm"
)

// AddOrder add a oder
func (order *Order) AddOrder(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(order)
	return tx.Error
}

// FetchOrder get specified order
func (order *Order) FetchOrder(ctx context.Context) ([]*Order, error) {
	orders := make([]*Order, 0)
	rows, err := ctx.Value("DbSession").(*gorm.DB).Table("orders").
		Where("contract = ? and direction = ? and deleted_at is null", order.Contract, order.Direction).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order := &Order{}
		if err := ctx.Value("DbSession").(*gorm.DB).ScanRows(rows, order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// DeleteOrder add a oder
func (order *Order) DeleteOrder(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Table("orders").
		Where("contract = ? and direction = ?", order.Contract, order.Direction).Delete(order)
	return tx.Error
}

// UpdateOrder update a oder
func (order *Order) UpdateOrder(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Table("orders").
		Where("contract = ? and direction = ? and deleted_at is null", order.Contract, order.Direction).Updates(order)
	return tx.Error
}
