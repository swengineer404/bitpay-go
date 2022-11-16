package bitpay

import "fmt"

type Client struct {
	api *apiClient
}

func New(key string, sandbox bool) *Client {
	url := "https://bitpay.com"
	if sandbox {
		url = "https://test.bitpay.com"
	}
	return &Client{
		api: newAPIClient(url, key),
	}
}

func (c *Client) CreateInvoice(params *CreateInvoiceParams) (*CreateInvoice, error) {
	var res CreateInvoice
	return &res, c.api.send("/invoices", "POST", params, &res)
}

func (c *Client) GetBTCAddress(invoiceId string) (string, error) {
	var res InvoiceData
	if err := c.api.send(fmt.Sprintf("/invoices/%s", invoiceId), "POST", nil, &res); err != nil {
		return "", err
	}

	return res.Invoice.PaymentCodes.BTC.Address, nil
}
