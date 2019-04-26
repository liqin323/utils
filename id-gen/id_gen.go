package idgen

import (
	"math/rand"
	"time"
	resSvc "utils/id-gen/resource"
	resModels "utils/id-gen/resource/models"
)

const (
	radixCharList = "QWERTYUPASDFGHJKLZXCVBNM023456789"
	radixNumList  = "0123456789"
)

func Initialize() (err error) {
	return resSvc.Initialize()
}

func genRandomCode(length int) (code string, err error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	codeGen := ""
	radixCharLength := len(radixCharList)
	for i := 0; i < length; i++ {
		codeGen = codeGen + string(radixCharList[r.Intn(radixCharLength)])
	}

	return codeGen, nil
}

func genRandomNumCode(length int) (code string, err error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	codeGen := ""
	numCharLength := len(radixNumList)
	for i := 0; i < length; i++ {
		codeGen = codeGen + string(radixNumList[r.Intn(numCharLength)])
	}

	return codeGen, nil
}

func NewUniqueID(idType string, length int) (code string, err error) {

	for true {
		code, err := genRandomCode(length)
		if err != nil {
			return "", err
		}

		typeCode := idType + code
		err = resSvc.UniqueCode.Get(idType, code)
		if err != nil {
			unqueCodeM := resModels.UniqueCode{
				Type: idType,
				Code: code,
			}
			resSvc.UniqueCode.Create(&unqueCodeM)
			return typeCode, nil
		}
	}

	return "", nil
}

func NewUniqueNumID(idType string, length int) (code string, err error) {

	for true {
		code, err := genRandomNumCode(length)

		if err != nil {
			return "", err
		}

		typeCode := idType + code
		err = resSvc.UniqueCode.Get(idType, code)
		if err != nil {
			unqueCodeM := resModels.UniqueCode{
				Type: idType,
				Code: code,
			}
			resSvc.UniqueCode.Create(&unqueCodeM)
			return typeCode, nil
		}
	}

	return "", nil
}
