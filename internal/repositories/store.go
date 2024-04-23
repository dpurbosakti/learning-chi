package store

import (
	account "github.com/dpurbosakti/go-native/internal/repositories/account/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	account.Querier
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*account.Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  account.New(connPool),
	}
}
