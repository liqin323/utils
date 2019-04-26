package profile

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Unknwon/goconfig"
)

var File *goconfig.ConfigFile

func LoadConfigFile() (err error) {
	// load profile
	svr_dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	prof_file := path.Join(svr_dir, "./.profile")

	File, err = goconfig.LoadConfigFile(prof_file)
	if err != nil {
		return
	}

	return
}

func LoadConfigFromFile(file string) (err error) {
	// load profile
	File, err = goconfig.LoadConfigFile(file)
	if err != nil {
		return
	}

	return
}
