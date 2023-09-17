package implement

import (
	"os"
	"strings"

	"github.com/chmikata/csvconvert/internal/domain/model"
	"github.com/chmikata/csvconvert/internal/userinterface"
	"github.com/chmikata/csvconvert/internal/utils/csv"
)

var _ userinterface.OutputWriter = (*FileWriter)(nil)

type FileWriter struct {
	filePath string
}

type Option func(*FileWriter)

func WithFilePath(filePath string) Option {
	return func(f *FileWriter) {
		f.filePath = filePath
	}
}

func NewFileWriter(options ...Option) *FileWriter {
	f := &FileWriter{
		filePath: "out.csv",
	}
	for _, option := range options {
		option(f)
	}
	return f
}

func (f *FileWriter) WriteCSV(data []model.HbEnvironment) error {
	csv, err := csv.Marshal(data)
	if err != nil {
		return err
	}
	file, err := os.Create(f.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range csv {
		file.WriteString(strings.Join(v, ",") + "\n")
	}
	return nil
}
