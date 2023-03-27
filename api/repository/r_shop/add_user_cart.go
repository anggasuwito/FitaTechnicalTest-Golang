package r_shop

import "FitaTechnicalTest-Golang/api/models"

func (r ShopRepo) AddtoUserCart(req models.CartInsertRequest) (err error) {
	stmt, err := r.DB.
		Prepare(`
			INSERT INTO cart
			(user_id,item_sku,qty)
				VALUES
			(?,?,?)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(req.UserID, req.ItemSKU, req.Quantity)
	return err
}
