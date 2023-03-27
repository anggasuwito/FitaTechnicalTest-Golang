package models

type Item struct {
	ID       string  `json:"id"`
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CartListResponse struct {
	CheckoutTotalPrice float64 `json:"checkout_total_price"`
	Items              []Cart  `json:"items"`
}

type Cart struct {
	CartID       string  `json:"cart_id"`
	UserID       string  `json:"user_id"`
	ItemSKU      string  `json:"item_sku"`
	ItemName     string  `json:"item_name"`
	ItemPrice    float64 `json:"item_price"`
	CartQuantity int     `json:"cart_quantity"`
}

type CartInsertRequest struct {
	UserID   string `json:"user_id"`
	ItemSKU  string `json:"item_sku"`
	Quantity int    `json:"quantity"`
}

type Transaction struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
