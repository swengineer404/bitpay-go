package bitpay

type CreateInvoiceParams struct {
	Description     string  `json:"itemDesc"`
	Price           float64 `json:"price"`
	Currency        string  `json:"currency"`
	OrderId         string  `json:"orderId"`
	NotificationURL string  `json:"notificationURL"`
	RedirectURL     string  `json:"redirectURL"`
	Buyer           Buyer   `json:"buyer"`
}

type CreateInvoice struct {
	Data struct {
		Invoice
		ExtendNotifications bool `json:"extendedNotifications"`
	} `json:"data"`
}

type InvoiceData struct {
	Invoice struct {
		Invoice
	} `json:"invoice"`
}

type InvoiceResult struct {
	Data struct {
		Invoice
	} `json:"data"`
}

type WebhookEvent struct {
	Event struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	Data struct {
		Invoice
	} `json:"data"`
}
