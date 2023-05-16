package bitpay

import "strconv"

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

const (
	InvoiceExceptionUnderpaid = "paidPartial"
	InvoiceExceptionOverpaid  = "paidOver"
)

type MinerFee struct {
	TotalFee int `json:"totalFee"`
}

func (m MinerFee) FormatTotalFee() string {
	return strconv.FormatFloat(float64(m.TotalFee)/float64(100000000), 'f', 6, 64)
}

type Invoice struct {
	ID                   string `json:"id"`
	OrderID              string `json:"orderId"`
	URL                  string `json:"url"`
	Status               string `json:"status"`
	ExceptionStatus      any    `json:"exceptionStatus"`
	PaymentDisplayTotals struct {
		BTC string `json:"BTC"`
	} `json:"paymentDisplayTotals"`
	PaymentDisplaySubTotals struct {
		BTC string `json:"BTC"`
	} `json:"paymentDisplaySubTotals"`
	MinerFees struct {
		BTC MinerFee `json:"BTC"`
	} `json:"minerFees"`
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
