package repository

import "github.com/chmikata/environment_csv/internal/domain/model"

type HbEnvironmentRepository interface {
	GetHbEnvironment() ([]model.HbEnvironment, error)
}
