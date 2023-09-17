package repository

import "github.com/chmikata/csvconvert/internal/domain/model"

type HbEnvironmentRepository interface {
	CreateEnvironment() ([]model.HbEnvironment, error)
}
