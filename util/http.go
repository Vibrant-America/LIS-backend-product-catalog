package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetResponse(method string, url string, body interface{}, sbody interface{}) (int, error) {
	bd, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bd))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res.StatusCode, err
	}

	sbd, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(sbd, sbody)
	if err != nil {
		return -1, err
	}
	return res.StatusCode, nil
}
