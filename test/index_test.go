package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func printResponse(t *testing.T, rsp *http.Response) {
	buff, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	t.Logf("status code: %d\nresponse data: %s\n", rsp.StatusCode, string(buff))
}

func Test_GetTableRecords(t *testing.T) {
	tableName := ""
	rsp, err := http.Get("http://127.0.0.1:8081/getTableRecords?tableName=" + tableName)
	if err != nil {
		t.Fatal(err)
	}
	printResponse(t, rsp)
}
