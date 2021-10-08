package utils

import (
	"encoding/json"

	"github.com/golang/glog"
)

func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		glog.Error("Failed to marshal data")
		return err
	}

	json.Unmarshal(data, dest)

	return nil
}
