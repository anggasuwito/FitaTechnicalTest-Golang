package r_shop

import "FitaTechnicalTest-Golang/api/models"

func (r ShopRepo) GetAllItem() (res []models.Item, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
				i.id,
				i.sku,
				i.name,
				i.price,
				i.qty
			FROM item i
			WHERE i.deleted_at IS NULL
			ORDER BY i.name ASC
		`)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err = rows.Scan(
			&item.ID,
			&item.SKU,
			&item.Name,
			&item.Price,
			&item.Quantity,
		); err != nil {
			return res, err
		}

		res = append(res, item)
	}
	return res, err
}
