package r_shop

import "FitaTechnicalTest-Golang/api/models"

func (r ShopRepo) UpdateUserCart(id string, req models.CartInsertRequest) (err error) {
	_, err = r.DB.
		Exec(`
			UPDATE
			    cart
			SET
			    qty = qty + ?
			WHERE id = ?
				AND deleted_at IS NULL
		`, req.Quantity, id)
	return err
}
