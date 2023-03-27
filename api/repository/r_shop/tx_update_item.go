package r_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r ShopRepo) TXUpdateItem(tx *sql.Tx, cart models.Cart) (err error) {
	_, err = tx.
		Exec(`
			UPDATE
			    item
			SET
			    qty = qty - ?
			WHERE sku = ?
				AND deleted_at IS NULL
		`, cart.CartQuantity, cart.ItemSKU)
	return err
}
