package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	file, err := os.Open("./_test/userCreate.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytesData, _ := ioutil.ReadAll(file)
	req, _ := http.NewRequest("POST", testServer.URL+"/users", bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json")
	res, err := testServer.Client().Do(req)
	if err != nil {
		t.Errorf("http request failed error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected to be get status ok but got not status:%v error: %s", res.Status, string(body))
	}
	body, err = ioutil.ReadAll(res.Body)
	t.Logf("reponse body :%s", string(body))
}
