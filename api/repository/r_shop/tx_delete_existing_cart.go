package r_shop

import (
	"database/sql"
)

func (r ShopRepo) TXDeleteExistingCart(tx *sql.Tx, userID string) (err error) {
	_, err = tx.
		Exec(`
			UPDATE
			    cart
			SET
			    deleted_at = now()
			WHERE user_id = ?
				AND deleted_at IS NULL
		`, userID)
	return err
}
