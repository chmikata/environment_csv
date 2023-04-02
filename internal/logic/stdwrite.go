package logic

import (
	"fmt"

	"github.com/chmikata/csvconvert/internal/convert"
)

var _ CSVWriter = (*StdOutWriter)(nil)

type StdOutWriter struct {
}

func NewStdOutWriter() *StdOutWriter {
	return &StdOutWriter{}
}

func (s *StdOutWriter) WriteCSV(data []CSV) error {
	csv, err := convert.Marshal(data)
	if err != nil {
		return err
	}
	for _, v := range csv {
		fmt.Println(v)
	}
	return nil
}
