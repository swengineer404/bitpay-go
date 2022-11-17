package bitpay

const (
	InvoiceStatusNew       = "new"
	InvoiceStatusPaid      = "paid"
	InvoiceStatusConfirmed = "confirmed"
	InvoiceStatusComplete  = "complete"
	InvoiceStatusInvalid   = "invalid"
	InvoiceStatusExpired   = "expired"
)

const (
	EventCodePaid            = 1003
	EventCodeExpired         = 1004
	EventCodeConfirmed       = 1005
	EventCodeCompleted       = 1006
	EventCodeFailedToConfirm = 1013
)

type Invoice struct {
	ID                   string `json:"id"`
	OrderID              string `json:"orderId"`
	URL                  string `json:"url"`
	Status               string `json:"status"`
	PaymentDisplayTotals struct {
		BTC string `json:"BTC"`
	} `json:"paymentDisplayTotals"`
	Buyer        Buyer
	PaymentCodes struct {
		BTC struct {
			Address string `json:"ADDRESS"`
		} `json:"BTC"`
	} `json:"paymentCodes"`
}

type Buyer struct {
	Email string `json:"email"`
}
