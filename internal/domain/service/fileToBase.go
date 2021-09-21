package service

import (
	"encoding/base64"
	"github.com/spf13/viper"
	"io/ioutil"
)

func fileToBase(fileName string, pathConfig string) *string {
	var b64 string
	path := viper.GetString("Files." + pathConfig)
	dir, err := ioutil.ReadFile(path + "/" + fileName)
	b64 = base64.StdEncoding.EncodeToString(dir)
	if err != nil {
		return nil
	}

	return &b64
}
