module github.com/domenicomastrangelo/safe

go 1.15

replace github.com/domenicomastrangelo/safe/database => ./database

replace github.com/domenicomastrangelo/safe/config => ./config

replace github.com/domenicomastrangelo/safe/encryption => ./encryption

replace github.com/domenicomastrangelo/safe/service => ./service

require github.com/mattn/go-sqlite3 v1.14.6
