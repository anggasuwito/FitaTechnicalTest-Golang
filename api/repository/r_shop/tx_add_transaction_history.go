package r_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r ShopRepo) TXAddTransactionHistory(tx *sql.Tx, req models.Cart, transactionID string) (err error) {
	stmt, err := tx.
		Prepare(`
			INSERT INTO transaction_history
			(transaction_id,user_id,cart_id)
				VALUES
			(?,?,?)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(transactionID, req.UserID, req.CartID)
	return err
}
