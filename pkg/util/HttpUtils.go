package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(method, uri string, body interface{}) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	var (
		res *http.Request
		err error
	)
	if body != nil {
		byteBody, _ := json.Marshal(body)
		res, err = http.NewRequest(method, uri, bytes.NewBuffer(byteBody))
	} else {
		res, err = http.NewRequest(method, uri, nil)
	}
	res.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	resp, ee := client.Do(res)
	if ee != nil {
		return "", ee
	}
	data, e := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("The request failure status code is %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	if e != nil {
		return "", e
	}
	bd := string(data)
	return bd, nil
}

func GetJsonVal(body string, key string) interface{} {
	data := gojsonq.New().JSONString(body)
	val := data.Find(key)
	if val == nil {
		return ""
	}
	return val
}
