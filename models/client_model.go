package models

import (
	"errors"
)

type Client struct {
	JsonModel
	Name        string  `json:"name"`
	TimeSeconds uint    `json:"timeSeconds"`
	PayRate     float32 `json:"rate"`
	Invoices    []Invoice
}

func CreateClient(name string, rate float32) (Client, error) {
	if (name == "") || (rate < 0) {
		return Client{}, errors.New("invalid client data")
	}

	client := Client{
		Name:        name,
		TimeSeconds: 0,
		PayRate:     rate,
	}

	return client, nil
}
