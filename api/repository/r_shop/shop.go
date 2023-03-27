package r_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"database/sql"
)

type ShopRepo struct {
	DB *sql.DB
}

type ShopRepoInterface interface {
	GetAllItem() (res []models.Item, err error)
	GetItemBySKU(itemSKU string) (res models.Item, err error)
	AddtoUserCart(req models.CartInsertRequest) (err error)
	UpdateUserCart(id string, req models.CartInsertRequest) (err error)
	GetUserCart(userID string) (res []models.Cart, err error)
	TXBegin() (tx *sql.Tx, err error)
	TXGetItemForUpdate(tx *sql.Tx, itemSKU string) (res models.Item, err error)
	TXUpdateItem(tx *sql.Tx, cart models.Cart) (err error)
	TXAddTransaction(tx *sql.Tx, req models.Transaction) (err error)
	TXAddTransactionHistory(tx *sql.Tx, req models.Cart, transactionID string) (err error)
	TXDeleteExistingCart(tx *sql.Tx, userID string) (err error)
}

func NewShopRepo(db *sql.DB) ShopRepoInterface {
	return &ShopRepo{DB: db}
}
