package resource

import "utils/id-gen/resource/models"

type I_UniqueCode interface {
	Create(code *models.UniqueCode) (err error)
	Get(codeType, code string) (err error)
	Delete(codeType, code string) (err error)
}
