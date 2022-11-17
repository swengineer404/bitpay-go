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

func (c *Client) GetInvoice(id string) (*Invoice, error) {
	var res InvoiceResult
	return &res.Data.Invoice, c.api.send("/invoices/"+id, "GET", nil, &res)
}

func (c *Client) GetBTCAddress(invoiceId string) (string, error) {
	var res InvoiceData
	if err := c.api.send(fmt.Sprintf("/invoiceData/%s", invoiceId), "POST", nil, &res); err != nil {
		return "", err
	}

	return res.Invoice.PaymentCodes.BTC.Address, nil
}
