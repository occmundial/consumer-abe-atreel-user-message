package interfaces

import "github.com/occmundial/consumer-abe-atreel-user-message/models"

type ITlaloc interface {
	GetLocTlaloc() (map[string]models.TlalocLocation, error)
}
