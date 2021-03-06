package domain

const (
	ACCT_ACTIVE = 3
	ACCT_LOCK   = 5
	ACCT_CLOSED = 6
)

type Account struct {
	AccountId   int     `db:"account_id"`
	CustomerId  int     `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Balance     float64 `db:"balance"`
	Status      int     `db:"status"`
}

type AccountRepository interface {
	Create(account *Account) (int, error)
	GetBalance(accountId int) (float64, error)
	Lock(accountId int) error
	Unlock(accountId int) error
}
