package logic

import ()

type CSV struct {
	Id      string `csv:"id"`
	Header1 string `csv:"header1"`
	Header2 string `csv:"header2"`
	Header3 string `csv:"header3"`
	Detail1 string `csv:"detail1"`
	Detail2 string `csv:"detail2"`
	Detail3 string `csv:"detail3"`
}

type CSVCreater interface {
	CreateCSV() ([]CSV, error)
}

type CSVWriter interface {
	WriteCSV(data []CSV) error
}

type Controller struct {
	csvCreater CSVCreater
	csvWriter  CSVWriter
}

func NewController(creater CSVCreater, writer CSVWriter) *Controller {
	return &Controller{
		csvCreater: creater,
		csvWriter:  writer,
	}
}

func (c *Controller) CreateCSV() error {
	data, err := c.csvCreater.CreateCSV()
	if err != nil {
		return err
	}
	if err := c.csvWriter.WriteCSV(data); err != nil {
		return err
	}
	return nil
}
