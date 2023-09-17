package infra

import (
	"database/sql"
	"fmt"

	"github.com/chmikata/csvconvert/internal/domain/model"
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

type HbEnvironmentRepository struct {
	host   string
	port   int
	user   string
	pass   string
	dbName string
}

type Option func(*HbEnvironmentRepository)

func WithUser(user string) Option {
	return func(repo *HbEnvironmentRepository) {
		repo.user = user
	}
}

func WithPass(pass string) Option {
	return func(repo *HbEnvironmentRepository) {
		repo.pass = pass
	}
}

func WithDbName(dbname string) Option {
	return func(repo *HbEnvironmentRepository) {
		repo.dbName = dbname
	}
}

func NewHbEnvironmentRepository(host string, port int, options ...Option) *HbEnvironmentRepository {
	repo := &HbEnvironmentRepository{
		host:   host,
		port:   port,
		user:   "postgres",
		pass:   "postgres",
		dbName: "postgres",
	}

	for _, option := range options {
		option(repo)
	}
	return repo
}

func (hr *HbEnvironmentRepository) CreateEnvironment() ([]model.HbEnvironment, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		hr.user, hr.pass, hr.dbName, hr.host, hr.port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	contracts, err := hr.searchContracts(db)
	if err != nil {
		return nil, err
	}
	var csv []model.HbEnvironment
	for _, contract := range contracts {
		details, err := hr.searchDetails(db, contract.Id)
		if err != nil {
			return nil, err
		}
		for _, detail := range details {
			csv = append(csv, model.HbEnvironment{
				Id:      contract.Id,
				Header1: contract.Header1,
				Header2: contract.Header2,
				Header3: contract.Header3,
				Detail1: detail.Detail1,
				Detail2: detail.Detail2,
				Detail3: detail.Detail3,
			})
		}
	}
	return csv, nil
}

func (hr *HbEnvironmentRepository) searchContracts(db *sql.DB) ([]Contract, error) {
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

func (hr *HbEnvironmentRepository) searchDetails(db *sql.DB, id string) ([]Detail, error) {
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
