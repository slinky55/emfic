package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	JsonModel
	TimeSeconds uint    `json:"timeSeconds"`
	AmtBilled   float32 `json:"amtBilled"`
	ClientID    uint    `json:"clientId"`
}

func CreateInvoice(c *Client, db *gorm.DB) (Invoice, error) {
	invoice := Invoice{
		TimeSeconds: c.TimeSeconds,
		AmtBilled:   (float32(c.TimeSeconds) / 3600) * c.PayRate,
	}

	res := db.Create(&invoice)

	if res.Error != nil {
		return Invoice{}, res.Error
	}

	return invoice, nil
}
