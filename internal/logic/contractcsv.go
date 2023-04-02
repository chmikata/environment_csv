package logic

import "github.com/chmikata/csvconvert/internal/infra"

type DataSearcher interface {
	SearchContracts() ([]infra.Contract, error)
	SearchDetails(id string) ([]infra.Detail, error)
}

var _ CSVCreater = (*ContractCSVCreater)(nil)

type ContractCSVCreater struct{}

func NewContractCSVCreater() *ContractCSVCreater {
	return &ContractCSVCreater{}
}

func (c *ContractCSVCreater) CreateCSV() ([]CSV, error) {
	sarcher := infra.NewPgDB("localhost", 5432,
		infra.WithUser("postgres"),
		infra.WithPass("postgres"),
		infra.WithDbName("convert"),
	)
	contracts, err := sarcher.SearchContracts()
	if err != nil {
		return nil, err
	}
	var csv []CSV
	for _, contract := range contracts {
		details, err := sarcher.SearchDetails(contract.Id)
		if err != nil {
			return nil, err
		}
		for _, detail := range details {
			csv = append(csv, CSV{
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
