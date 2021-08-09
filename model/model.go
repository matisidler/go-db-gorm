package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string         `gorm:"type:varchar(100);not null"`
	Observaciones sql.NullString `gorm:"type:varchar(100);"`
	Price         int            `gorm:"not null"`
	InvoiceItems  []InvoiceItem
}

type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}

type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100);not null"`
	InvoiceItems []InvoiceItem
}
