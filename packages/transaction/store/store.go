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
	GetTransactionById(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*Transaction, error)
	GetTotalSales(ctx context.Context) (float32, error)
	GetSalesByProductId(ctx context.Context, productId uuid.UUID) (float32, error)
	GetTopFiveCustomersId(ctx context.Context) ([]*uuid.UUID, error)
}

type store struct {
	DB database.Database
}

func NewTransactionStore(db database.Database) Store {
	return &store{DB: db}
}

func (s store) GetTotalSales(ctx context.Context) (float32, error) {
	totalPrice := 0
	q := `
		SELECT SUM(total_price)
		FROM "transaction"
		`
	row := s.DB.QueryRow(ctx, q)

	err := row.Scan(
		&totalPrice,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return 0, nil
}

func (s store) GetSalesByProductId(ctx context.Context, productId uuid.UUID) (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetTopFiveCustomersId(ctx context.Context) ([]*uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) InsertTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
	log.Println("Store: Creating new Transaction...")

	createdAt := time.Now().UTC()
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

func (s store) GetTransactionById(ctx context.Context, id uuid.UUID) (*Transaction, error) {

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
			return nil, nil
		}
		return nil, err
	}

	return transaction, nil
}

func (s store) GetAllTransaction(ctx context.Context) ([]*Transaction, error) {
	q := `SELECT 
            id,
            created_at,
            customer_id,
            product_id,
            quantity,
            total_price 
        FROM transaction order by created_at`

	rows, err := s.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []*Transaction

	for rows.Next() {
		transaction := &Transaction{}
		err := rows.Scan(
			&transaction.Id,
			&transaction.CreatedAt,
			&transaction.CustomerId,
			&transaction.ProductId,
			&transaction.Quantity,
			&transaction.TotalPrice,
		)
		transactions = append(transactions, transaction)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
