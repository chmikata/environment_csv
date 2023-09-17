package application

import (
	"github.com/chmikata/environment_csv/internal/domain/repository"
	"github.com/chmikata/environment_csv/internal/userinterface"
)

type Controller struct {
	repository repository.HbEnvironmentRepository
	writer     userinterface.OutputWriter
}

func NewController(repository repository.HbEnvironmentRepository,
	writer userinterface.OutputWriter) *Controller {
	return &Controller{
		repository: repository,
		writer:     writer,
	}
}

func (c *Controller) CreateEnvironment() error {
	data, err := c.repository.GetHbEnvironment()
	if err != nil {
		return err
	}
	if err := c.writer.WriteCSV(data); err != nil {
		return err
	}
	return nil
}
