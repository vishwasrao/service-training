package schema

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

const seeds = `
INSERT INTO products (product_id, name, cost, quantity, date_created, date_updated) VALUES
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'Comic Books', 50, 42, '2019-01-01 00:00:01.000001+00', '2019-01-01 00:00:01.000001+00'),
	('72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'McDonalds Toys', 75, 120, '2019-01-01 00:00:02.000001+00', '2019-01-01 00:00:02.000001+00');

INSERT INTO sales (sale_id, product_id, quantity, paid) VALUES
	('98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', 'a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 2, 100),
	('85f6fb09-eb05-4874-ae39-82d1a30fe0d7', 'a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 5, 250),
	('a235be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 3, 225);
`
