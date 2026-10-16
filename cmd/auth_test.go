package main

import (
	"bytes"
	"encoding/json"
	"http-server/internal/auth"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "dev@dev.com",
		Password: "qwerty",
	})

	response, err := http.Post(
		ts.URL+"/auth/login",
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatalf("err: token empty")
	}
}

func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "dev@dev.com",
		Password: "qwerty",
	})

	response, err := http.Post(
		ts.URL+"/auth/login",
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected: %d, got: %d", http.StatusUnauthorized, response.StatusCode)
	}
}
