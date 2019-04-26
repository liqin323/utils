package resource

import (
	mongoImpl "utils/id-gen/resource/impl/mongo"
)

var UniqueCode I_UniqueCode

func Initialize() (err error) {

	// mongo db initialize
	_err := mongoImpl.Initialize()
	if _err != nil {
		err = _err
		return
	}

	UniqueCode = new(mongoImpl.UniqueCodeDAO)
	return
}
