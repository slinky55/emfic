package models

type Invoice struct {
	JsonModel
	TimeSeconds uint    `json:"timeSeconds"`
	AmtBilled   float32 `json:"amtBilled"`
	ClientID    uint    `json:"clientId"`
}

func CreateInvoice(c *Client) (Invoice, error) {

	invoice := Invoice{
		TimeSeconds: c.TimeSeconds,
		AmtBilled:   (float32(c.TimeSeconds) / 3600) * c.PayRate,
	}

	return invoice, nil
}
