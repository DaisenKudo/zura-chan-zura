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
	go func() {
		main()
	}()

	time.Sleep(time.Second * 5)

	result, err := http.Get("http://localhost:" + strconv.Itoa(port))
	if err != nil {
		//t.Fatal("サーバーが起動しませんでした")
	}
	if result.StatusCode != 200 {
		t.Fatal("サーバーの設定に問題があります\n", result.Header)
	}
}
