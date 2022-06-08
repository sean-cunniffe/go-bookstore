package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseBody(r io.Reader, x interface{}) error {
	if body, err := ioutil.ReadAll(r); err == nil {
		if err = json.Unmarshal(body, &x); err != nil {
			return err
		}
	}
	return nil
}
