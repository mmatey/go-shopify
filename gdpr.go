package shopify

// GDPRCustomer contains the customer's information for the redaction.
type GDPRCustomer struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// CustomersRedact is the unmarshal struct for the `customers/redact` webhook.
type GDPRCustomersRedact struct {
	ShopID     int64         `json:"shop_id"`
	ShopDomain string        `json:"shop_domain"`
	Customer   *GDPRCustomer `json:"customer"`
	OrderIDs   []int64       `json:"orders_to_redact"`
}

// GDPRShopRedact is the unmarshal struct for the `shop/redact` webhook.
type GDPRShopRedact struct {
	ShopID     int64  `json:"shop_id"`
	ShopDomain string `json:"shop_domain"`
}

// GDPRCustomersRequest is the unmarshal struct for the `customers/data_request` webhook.
type GDPRCustomersRequest struct {
	ShopID     int64         `json:"shop_id"`
	ShopDomain string        `json:"shop_domain"`
	Customer   *GDPRCustomer `json:"customer"`
	OrderIDs   []int64       `json:"orders_requested"`
}
