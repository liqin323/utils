package mongoImpl

import (
	"errors"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"utils/profile"
	"utils/slog"
)

var mgo_session *mgo.Session

type dbInfo struct {
	Host   []string
	UName  string
	Psw    string
	DBName string
}

var DBInfo dbInfo

func Initialize() (err error) {

	Host, _err := profile.File.GetValue("Database", "host")
	if _err != nil {
		err = _err
		return
	}

	DBInfo.Host = strings.Split(Host, ",")

	DBInfo.UName, _ = profile.File.GetValue("Database", "user_name")
	DBInfo.Psw, _ = profile.File.GetValue("Database", "pass_word")

	DBInfo.DBName, _err = profile.File.GetValue("Database", "db_name")
	if _err != nil {
		err = _err
		return
	}

	return
}

func InitializeEx(host, user, psw, dbname string) (err error) {

	DBInfo.Host = strings.Split(host, ",")

	DBInfo.UName = user
	DBInfo.Psw = psw

	DBInfo.DBName = dbname

	return
}

func Finalize() {
	if mgo_session != nil {
		mgo_session.Close()
		mgo_session = nil
	}
}

func Session() *mgo.Session {

	if mgo_session == nil {
		var dialInfo mgo.DialInfo

		dialInfo.Database = DBInfo.DBName
		dialInfo.Addrs = DBInfo.Host

		if DBInfo.UName != "" && DBInfo.Psw != "" {
			dialInfo.Username = DBInfo.UName
			dialInfo.Password = DBInfo.Psw
		}

		err := errors.New("")
		mgo_session, err = mgo.DialWithInfo(&dialInfo)
		if err != nil || mgo_session == nil {
			slog.Error("fail to dail with DB, err:%v, DB info:%v", err, DBInfo)
		}

		mgo_session.SetMode(mgo.Strong, true)
	}

	return mgo_session.Clone()
}

func Collection(c_name string, session *mgo.Session) (c *mgo.Collection) {
	db := session.DB(DBInfo.DBName)
	c = db.C(c_name)
	return
}

func CollectionEx(c_name, DBName string, session *mgo.Session) (c *mgo.Collection) {
	db := session.DB(DBName)
	c = db.C(c_name)
	return
}
