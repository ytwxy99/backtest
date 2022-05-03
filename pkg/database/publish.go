package database

import (
	"context"

	"gorm.io/gorm"
)

// AddPublish add a publish record
func (publish *Publish) AddPublish(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Create(publish)
	return tx.Error
}

// DeletePublish add a publish record
func (publish *Publish) DeletePublish(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Delete(publish)
	return tx.Error
}

// UpdatePublish add a publish record
func (publish *Publish) UpdatePublish(ctx context.Context) error {
	tx := ctx.Value("DbSession").(*gorm.DB).Updates(publish)
	return tx.Error
}

// GetAllPublishes get all publish
func GetAllPublishes(ctx context.Context) ([]Publish, error) {
	var publishes []Publish
	tx := ctx.Value("DbSession").(*gorm.DB).Find(&publishes)
	return publishes, tx.Error
}
