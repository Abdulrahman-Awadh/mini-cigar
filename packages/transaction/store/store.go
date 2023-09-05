package store

import (
	"ecommerece/packages/database"
	"errors"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Transaction struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	CustomerId uuid.UUID
	ProductId  uuid.UUID
	Quantity   int32
	TotalPrice float32
}

type Store interface {
	InsertTransaction(ctx context.Context, transaction Transaction) (*Transaction, error)
	GetTransactionById(ctx context.Context, id uuid.UUID) *Transaction
	GetAllTransactions(ctx context.Context) ([]*Transaction, error)
}
type store struct {
	DB database.Database
}

type TransactionStore struct {
	database.Database
}

func NewTransactionStore(db database.Database) Store {
	return &store{DB: db}
}
func (s store) InsertTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
	log.Println("Creating new Transaction...")

	createdAt := time.Now()
	id := uuid.New()

	insertSql := `INSERT INTO transaction (
			id,
			created_at,
			customer_id,
			product_id,
			quantity,
			total_price
		)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.DB.Exec(ctx, insertSql, id, createdAt, transaction.CustomerId, transaction.ProductId, transaction.Quantity, transaction.TotalPrice)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		Id:         id,
		CreatedAt:  createdAt,
		CustomerId: transaction.CustomerId,
		ProductId:  transaction.ProductId,
		Quantity:   transaction.Quantity,
		TotalPrice: transaction.TotalPrice,
	}, nil

}
func (s store) GetTransactionById(ctx context.Context, id uuid.UUID) *Transaction {

	q := `SELECT 
			id,
			created_at,
			customer_id,
			product_id,
			quantity,
			total_price 
		FROM transaction
		WHERE id = $1`

	row := s.DB.QueryRow(ctx, q, id)

	transaction := &Transaction{}
	err := row.Scan(
		&transaction.Id,
		&transaction.CreatedAt,
		&transaction.CustomerId,
		&transaction.ProductId,
		&transaction.Quantity,
		&transaction.TotalPrice,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return nil
	}

	return transaction
}
func (s store) GetAllTransactions(ctx context.Context) ([]*Transaction, error) {
	q := `SELECT 
            id,
            created_at,
            customer_id,
            product_id,
            quantity,
            total_price 
        FROM transaction`

	rows, err := s.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.CreatedAt,
			&transaction.CustomerId,
			&transaction.ProductId,
			&transaction.Quantity,
			&transaction.TotalPrice,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
