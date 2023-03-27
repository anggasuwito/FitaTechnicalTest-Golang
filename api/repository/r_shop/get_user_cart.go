package r_shop

import "FitaTechnicalTest-Golang/api/models"

func (r ShopRepo) GetUserCart(userID string) (res []models.Cart, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
				c.id,
				c.user_id,
				c.item_sku,
				i.name,
				i.price,
				c.qty
			FROM cart c
			LEFT JOIN item i ON i.sku = c.item_sku
			WHERE c.deleted_at IS NULL
				AND c.user_id = ?
			ORDER BY i.name
		`, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var cart models.Cart
		if err = rows.Scan(
			&cart.CartID,
			&cart.UserID,
			&cart.ItemSKU,
			&cart.ItemName,
			&cart.ItemPrice,
			&cart.CartQuantity,
		); err != nil {
			return res, err
		}

		res = append(res, cart)
	}
	return res, err
}
