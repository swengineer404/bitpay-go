package bitpay

type CreateInvoiceParams struct {
	Price           float64     `json:"price"`
	Currency        string      `json:"currency"`
	OrderId         string      `json:"order_id"`
	NotificationURL string      `json:"notification_url"`
	RedirectURL     string      `json:"redirect_url"`
	Buyer           BuyerParams `json:"buyer"`
}

type BuyerParams struct {
	Email string `json:"email"`
}

type CreateInvoice struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type InvoiceData struct {
	Invoice struct {
		PaymentCodes struct {
			BTC struct {
				Address string `json:"ADDRESS"`
			} `json:"BTC"`
		} `json:"paymentCodes"`
	} `json:"invoice"`
}
