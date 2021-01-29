module github.com/domenicomastrangelo/safe

go 1.15

replace github.com/domenicomastrangelo/safe/backup => ./backup

replace github.com/domenicomastrangelo/safe/database => ./database

replace github.com/domenicomastrangelo/safe/config => ./config

replace github.com/domenicomastrangelo/safe/service => ./service

replace github.com/domenicomastrangelo/safe/encryption => ./encryption

require github.com/mattn/go-sqlite3 v1.14.6

require (
	github.com/domenicomastrangelo/safe/backup v0.0.0
	github.com/domenicomastrangelo/safe/config v0.0.0
	github.com/domenicomastrangelo/safe/database v0.0.0-00010101000000-000000000000
	github.com/domenicomastrangelo/safe/encryption v0.0.0
	github.com/domenicomastrangelo/safe/service v0.0.0
)
