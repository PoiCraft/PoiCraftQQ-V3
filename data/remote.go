package data

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendCommand(command string) (string, error) {
	body := bytes.NewBuffer([]byte(command))
	resp, err := http.Post(Conf.RemoteURL+"/execute_command", "text/plain", body)
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(result), nil
}
