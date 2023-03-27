package r_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r ShopRepo) TXGetItemForUpdate(tx *sql.Tx, itemSKU string) (res models.Item, err error) {
	err = tx.
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
			FOR UPDATE
		`, itemSKU).Scan(
		&res.ID,
		&res.SKU,
		&res.Name,
		&res.Price,
		&res.Quantity,
	)
	return res, err
}
