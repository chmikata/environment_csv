package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/chmikata/environment_csv/internal/domain/model"
	"github.com/chmikata/environment_csv/internal/domain/repository"
	"github.com/chmikata/environment_csv/internal/infra/boilerEntity"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type HbEnvironment struct {
	host   string
	port   int
	user   string
	pass   string
	dbName string
}

var _ repository.HbEnvironmentRepository = (*HbEnvironment)(nil)

type Option func(*HbEnvironment)

func WithUser(user string) Option {
	return func(he *HbEnvironment) {
		he.user = user
	}
}

func WithPass(pass string) Option {
	return func(he *HbEnvironment) {
		he.pass = pass
	}
}

func WithDbName(dbname string) Option {
	return func(he *HbEnvironment) {
		he.dbName = dbname
	}
}

func NewHbEnvironment(host string, port int, options ...Option) *HbEnvironment {
	he := &HbEnvironment{
		host:   host,
		port:   port,
		user:   "postgres",
		pass:   "postgres",
		dbName: "postgres",
	}

	for _, option := range options {
		option(he)
	}
	return he
}

func (he *HbEnvironment) GetHbEnvironment() ([]model.HbEnvironment, error) {
	db, err := he.initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	contracts, err := he.searchContracts(db)
	if err != nil {
		return nil, err
	}
	var heEnvironment []model.HbEnvironment
	for _, contract := range contracts {
		details, err := he.searchDetails(db, contract.ID)
		if err != nil {
			return nil, err
		}
		for _, detail := range details {
			heEnvironment = append(heEnvironment, model.HbEnvironment{
				Id:      contract.ID,
				Header1: contract.Header1.String,
				Header2: contract.Header2.String,
				Header3: contract.Header3.String,
				Detail1: detail.Detail1.String,
				Detail2: detail.Detail2.String,
				Detail3: detail.Detail3.String,
			})
		}
	}
	return heEnvironment, nil
}

func (he *HbEnvironment) initDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		he.user, he.pass, he.dbName, he.host, he.port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (he *HbEnvironment) searchContracts(db *sql.DB) ([]boilerEntity.Contract, error) {
	rows, err := boilerEntity.Contracts(
		qm.OrderBy("id ASC"),
	).All(context.Background(), db)
	if err != nil {
		return nil, err
	}

	var contracts []boilerEntity.Contract
	for _, row := range rows {
		contracts = append(contracts, *row)
	}
	return contracts, nil
}

func (he *HbEnvironment) searchDetails(db *sql.DB, id string) ([]boilerEntity.Detail, error) {
	rows, err := boilerEntity.Details(
		qm.Where("id = ?", id),
		qm.OrderBy("id ASC"),
	).All(context.Background(), db)
	if err != nil {
		return nil, err
	}

	var details []boilerEntity.Detail
	for _, row := range rows {
		details = append(details, *row)
	}
	return details, nil
}
