package logic

import (
	"github.com/chmikata/csvconvert/internal/domain/repository"
	"github.com/chmikata/csvconvert/internal/userinterface"
)

type HbEnvironmentController struct {
	repository repository.HbEnvironmentRepository
	writer     userinterface.OutputWriter
}

func NewController(repository repository.HbEnvironmentRepository,
	writer userinterface.OutputWriter) *HbEnvironmentController {
	return &HbEnvironmentController{
		repository: repository,
		writer:     writer,
	}
}

func (hc *HbEnvironmentController) CreateEnvironment() error {
	data, err := hc.repository.CreateEnvironment()
	if err != nil {
		return err
	}
	if err := hc.writer.WriteCSV(data); err != nil {
		return err
	}
	return nil
}
