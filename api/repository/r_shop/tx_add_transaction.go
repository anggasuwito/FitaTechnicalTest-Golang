package r_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"database/sql"
)

func (r ShopRepo) TXAddTransaction(tx *sql.Tx, req models.Transaction) (err error) {
	stmt, err := tx.
		Prepare(`
			INSERT INTO transaction
			(id,user_id,total_price,status)
				VALUES
			(?,?,?,?)
		`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(req.ID, req.UserID, req.TotalPrice, req.Status)
	return err
}
