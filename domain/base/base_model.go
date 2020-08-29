package base

import "time"

// DefaultModel a basic struct which includes following fields: Id, CreatedAt, UpdatedAt
type DefaultModel struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
