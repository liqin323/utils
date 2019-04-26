package ejson

import (
	"encoding/json"
	"utils/slog"
)

func Marshal(v interface{}) (res []byte, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		slog.Debug("Server Marshal error")
		return data, err
	}
	return Encrypt(data)
}

func Unmarshal(data []byte, v interface{}) (err error) {
	data, err = Decrypt(data)
	if err != nil {
		return err
	}
	//	return json.Unmarshal(data, v)
	err = json.Unmarshal(data, v)
	if err != nil {
		slog.Debug("Server Unmarshal error")
	}
	return err
}
