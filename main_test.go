package main

import (
	"net/http"
	"strconv"
	"testing"
	"time"
)

const (
	port = 8080
)

func TestMainSuccess(t *testing.T) {
	go main()

	time.Sleep(5000)

	result, err := http.Get("http://localhost:" + strconv.Itoa(port))
	if err != nil {
		t.Fatal("サーバーが起動しませんでした")
	}
	if result.StatusCode != 200 {
		t.Fatal("サーバーの設定に問題があります\n", result.Header)
	}
}
