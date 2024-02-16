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
	client := &http.Client{Timeout: 3 * time.Second}
	byteBody, _ := json.Marshal(body)
	res, err := http.NewRequest(method, uri, bytes.NewBuffer(byteBody))
	res.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	resp, ee := client.Do(res)
	data, e := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return "", ee
	}
	defer resp.Body.Close()
	if e != nil {
		return "", e
	}
	return string(data), nil
}

func GetJsonVal(body string, key string) interface{} {
	data := gojsonq.New().JSONString(body)
	val := data.Find(key)
	if val == nil {
		return ""
	}
	return val
}
