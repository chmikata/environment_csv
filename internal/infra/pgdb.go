package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Contract struct {
	Id      string
	Header1 string
	Header2 string
	Header3 string
}

type Detail struct {
	Id      string
	Detail1 string
	Detail2 string
	Detail3 string
}

type PgDB struct {
	host   string
	port   int
	user   string
	pass   string
	dbName string
}

type Option func(*PgDB)

func WithUser(user string) Option {
	return func(pgdb *PgDB) {
		pgdb.user = user
	}
}

func WithPass(pass string) Option {
	return func(pgdb *PgDB) {
		pgdb.pass = pass
	}
}

func WithDbName(dbname string) Option {
	return func(pgdb *PgDB) {
		pgdb.dbName = dbname
	}
}

func NewPgDB(host string, port int, options ...Option) *PgDB {
	pgdb := &PgDB{
		host:   host,
		port:   port,
		user:   "postgres",
		pass:   "postgres",
		dbName: "postgres",
	}

	for _, option := range options {
		option(pgdb)
	}
	return pgdb
}

func (p *PgDB) SearchContracts() ([]Contract, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		p.user, p.pass, p.dbName, p.host, p.port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM contract")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []Contract
	for rows.Next() {
		var contract Contract
		err := rows.Scan(&contract.Id, &contract.Header1, &contract.Header2, &contract.Header3)
		if err != nil {
			panic(err)
		}
		contracts = append(contracts, contract)
	}
	return contracts, nil
}

func (p *PgDB) SearchDetails(id string) ([]Detail, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		p.user, p.pass, p.dbName, p.host, p.port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM detail WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []Detail
	for rows.Next() {
		var detail Detail
		err := rows.Scan(&detail.Id, &detail.Detail1, &detail.Detail2, &detail.Detail3)
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}
	return details, nil
}
