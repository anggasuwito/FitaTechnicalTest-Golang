package r_shop

import (
	"database/sql"
)

func (r ShopRepo) TXBegin() (tx *sql.Tx, err error) {
	return r.DB.Begin()
}
