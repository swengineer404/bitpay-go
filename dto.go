package bitpay

type CreateInvoiceParams struct {
	Description         string  `json:"itemDesc"`
	Price               float64 `json:"price"`
	Currency            string  `json:"currency"`
	OrderId             string  `json:"orderId"`
	NotificationURL     string  `json:"notificationURL"`
	RedirectURL         string  `json:"redirectURL"`
	ExtendNotifications bool    `json:"extendedNotifications"`
	Buyer               Buyer   `json:"buyer"`
}

type CreateInvoice struct {
	Data struct {
		Invoice
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
	} `json:"event"`
	Data struct {
		Invoice
	} `json:"data"`
}
