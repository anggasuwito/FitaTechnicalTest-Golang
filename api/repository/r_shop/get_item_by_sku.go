package r_shop

import "FitaTechnicalTest-Golang/api/models"

func (r ShopRepo) GetItemBySKU(itemSKU string) (res models.Item, err error) {
	err = r.DB.
		QueryRow(`
			SELECT 
				i.id,
				i.sku,
				i.name,
				i.price,
				i.qty
			FROM item i
			WHERE i.sku = ?
				AND i.deleted_at IS NULL
		`, itemSKU).Scan(
		&res.ID,
		&res.SKU,
		&res.Name,
		&res.Price,
		&res.Quantity,
	)
	return res, err
}
