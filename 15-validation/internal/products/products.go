package products

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Predefined errors identify expected failure conditions.
var (
	// ErrNotFound is used when a specific Product is requested but does not exist.
	ErrNotFound = errors.New("product not found")
)

// Product is an item we sell.
type Product struct {
	ID       string `db:"product_id" json:"id"`
	Name     string `db:"name" json:"name" validate:"required"`
	Cost     int    `db:"cost" json:"cost" validate:"gte=0"`
	Quantity int    `db:"quantity" json:"quantity" validate:"gte=1"`
	Sold     int    `db:"sold" json:"sold"`
	Revenue  int    `db:"revenue" json:"revenue"`
}

// List gets all Products from the database.
func List(ctx context.Context, db *sqlx.DB) ([]Product, error) {
	var products []Product
	const q = `SELECT
	p.*,
	COALESCE(SUM(s.quantity) ,0) AS sold,
	COALESCE(SUM(s.paid), 0) AS revenue
FROM products AS p
LEFT JOIN sales AS s ON p.product_id = s.product_id
GROUP BY p.product_id`

	if err := db.SelectContext(ctx, &products, q); err != nil {
		return nil, errors.Wrap(err, "selecting products")
	}

	return products, nil
}

// Create uses the provided *Product to insert a new product record. The ID
// field provided is populated.
func Create(ctx context.Context, db *sqlx.DB, p *Product) error {
	p.ID = uuid.New().String()

	_, err := db.ExecContext(ctx, `
		INSERT INTO products
		(product_id, name, cost, quantity)
		VALUES ($1, $2, $3, $4)`,
		p.ID, p.Name, p.Cost, p.Quantity,
	)
	if err != nil {
		return errors.Wrap(err, "inserting product")
	}

	return nil
}

// Get finds the product identified by a given ID.
func Get(ctx context.Context, db *sqlx.DB, id string) (*Product, error) {
	var p Product

	const q = `SELECT
	p.*,
	COALESCE(SUM(s.quantity) ,0) AS sold,
	COALESCE(SUM(s.paid), 0) AS revenue
FROM products AS p
LEFT JOIN sales AS s ON p.product_id = s.product_id
WHERE p.product_id = $1
GROUP BY p.product_id`

	if err := db.GetContext(ctx, &p, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, errors.Wrap(err, "selecting single product")
	}

	return &p, nil
}
