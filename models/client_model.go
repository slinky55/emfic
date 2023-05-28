package models

import "gorm.io/gorm"

type Client struct {
	JsonModel
	Name        string  `json:"name"`
	TimeSeconds uint    `json:"timeSeconds"`
	PayRate     float32 `json:"rate"`
	Invoices    []Invoice
}

func CreateClient(name string, rate float32, db *gorm.DB) (Client, error) {
	client := Client{
		Name:        name,
		TimeSeconds: 0,
		PayRate:     rate,
	}

	res := db.Create(&client)

	if res.Error != nil {
		return Client{}, res.Error
	}

	return client, nil
}
