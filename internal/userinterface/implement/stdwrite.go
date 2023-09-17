package implement

import (
	"fmt"

	"github.com/chmikata/environment_csv/internal/domain/model"
	"github.com/chmikata/environment_csv/internal/userinterface"
	"github.com/chmikata/environment_csv/internal/utils/csv"
)

var _ userinterface.OutputWriter = (*StdOutWriter)(nil)

type StdOutWriter struct {
}

func NewStdOutWriter() *StdOutWriter {
	return &StdOutWriter{}
}

func (s *StdOutWriter) WriteCSV(data []model.HbEnvironment) error {
	csv, err := csv.Marshal(data)
	if err != nil {
		return err
	}
	for _, v := range csv {
		fmt.Println(v)
	}
	return nil
}
