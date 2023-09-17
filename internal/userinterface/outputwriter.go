package userinterface

import "github.com/chmikata/csvconvert/internal/domain/model"

type OutputWriter interface {
	WriteCSV(data []model.HbEnvironment) error
}
