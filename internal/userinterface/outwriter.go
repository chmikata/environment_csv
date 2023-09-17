package userinterface

import "github.com/chmikata/environment_csv/internal/domain/model"

type OutputWriter interface {
	WriteCSV(data []model.HbEnvironment) error
}
