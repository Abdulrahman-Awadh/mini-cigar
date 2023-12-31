package store

import (
	"ecommerece/packages/database"
	"errors"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"time"
)

type Transaction struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	CustomerId uuid.UUID
	ProductId  uuid.UUID
	Quantity   int32
	TotalPrice int64
}

type Store interface {
	InsertTransaction(ctx context.Context, transaction Transaction) (*Transaction, error)
	GetTransactionById(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*Transaction, error)
	GetTotalSales(ctx context.Context) (*int64, error)
	GetSalesByProductId(ctx context.Context, productId uuid.UUID) (*int64, error)
	GetTopFiveCustomersId(ctx context.Context) ([]*uuid.UUID, error)
}

type store struct {
	DB database.Database
}

func NewTransactionStore(db database.Database) Store {
	return &store{DB: db}
}

func (s store) GetTotalSales(ctx context.Context) (*int64, error) {
	var totalPrice *int64
	q := `
		SELECT SUM(total_price)
		FROM transaction
		`
	row := s.DB.QueryRow(ctx, q)
	err := row.Scan(
		&totalPrice,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return totalPrice, nil
}

func (s store) GetSalesByProductId(ctx context.Context, productId uuid.UUID) (*int64, error) {
	var totalPrice *int64
	q := `
			SELECT SUM(total_price)
			FROM "transaction"
			WHERE product_id = $1;
		`

	row := s.DB.QueryRow(ctx, q, productId)
	err := row.Scan(&totalPrice)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return totalPrice, nil
}

func (s store) GetTopFiveCustomersId(ctx context.Context) ([]*uuid.UUID, error) {
	var totalPrice int64
	var customerId uuid.UUID
	q := `
			select
				customer_id ,
				SUM (total_price)
			FROM
				"transaction" 
			GROUP BY
				customer_id  
			ORDER BY
				SUM (total_price) DESC
			limit 5;
			`

	rows, err := s.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customerIds []*uuid.UUID

	for rows.Next() {
		err := rows.Scan(
			&customerId,
			&totalPrice,
		)
		customerIds = append(customerIds, &customerId)
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

	return customerIds, nil

}

func (s store) InsertTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
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
