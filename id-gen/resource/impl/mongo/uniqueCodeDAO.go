package mongoImpl

import (
	"utils/id-gen/resource/models"

	"gopkg.in/mgo.v2/bson"
)

type UniqueCodeDAO struct {
}

var c_uniqueCode string = "unique.code"

func (self *UniqueCodeDAO) Create(code *models.UniqueCode) (err error) {

	s := Session()
	defer s.Close()

	c := Collection(c_uniqueCode, s)

	err = c.Insert(code)
	if err != nil {
		return
	}

	return
}

func (self *UniqueCodeDAO) Get(codeType, code string) (err error) {

	s := Session()
	defer s.Close()

	c := Collection(c_uniqueCode, s)

	_uniqueCode := new(models.UniqueCode)
	err = c.Find(bson.M{"code": code, "type": codeType}).One(_uniqueCode)
	if err != nil {
		return
	}

	return
}

func (self *UniqueCodeDAO) Delete(codeTpye, code string) (err error) {

	s := Session()
	defer s.Close()

	c := Collection(c_uniqueCode, s)

	err = c.Remove(bson.M{"code": code, "type": codeTpye})
	if err != nil {
		return
	}

	return
}
