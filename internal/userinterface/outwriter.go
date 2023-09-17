package userinterface

import "github.com/chmikata/environment_csv/internal/domain/model"

type OutputWriter interface {
	WriteCSV(hbEnviroment []model.HbEnvironment) error
}
