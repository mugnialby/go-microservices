package entities

import (
	"time"
)

/**
 * created by	: albym
 * created at	: 4 Jun 2024
 *
 * Users Entity Class
 */
type Users struct {
	ID         uint64 `gorm:"primarykey"`
	Username   string `gorm:"unique"`
	Password   string
	FirstName  string
	LastName   string
	Email      string `gorm:"unique"`
	ManagerId  uint64
	Status     string
	CreatedBy  string
	CreatedAt  time.Time
	ModifiedBy string
	ModifiedAt time.Time
}
