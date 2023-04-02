package logic

import (
	"os"
	"strings"

	"github.com/chmikata/csvconvert/internal/convert"
)

var _ CSVWriter = (*FileWriter)(nil)

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

func (f *FileWriter) WriteCSV(data []CSV) error {
	csv, err := convert.Marshal(data)
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
