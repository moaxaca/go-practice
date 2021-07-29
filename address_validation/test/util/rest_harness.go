package util

import (
	"io"
	"io.parcely.address_validation/internal"
	"net"
	"net/http"
)

type TestRestHarness struct {
	Hostname string
}

func (th *TestRestHarness) Get(path string) (resp *http.Response, err error) {
	return http.Get("http://localhost:8079"+path)
}

func (th *TestRestHarness) Post(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return http.Post("http://localhost:8079"+path, contentType, body)
}

func CreateTestRestHarness() TestRestHarness {
	port := "8079"
	config := internal.RestServerConfiguration{}
	config.Name = "test suite"
	config.Address = ":"+port
	th := TestRestHarness{}
	th.Hostname = "localhost"+config.Address
	// Test Port
	ln, errTest := net.Listen("tcp", ":" + port)
	if errTest != nil {
		return th
	}
	errClose := ln.Close()
	if errClose != nil {
		panic("Unable to close port")
	}
	// Start Server
	srv := internal.CreateRestServer(config)
	err := srv.Start()
	if err != nil {
		panic("Unable to create test server")
	}
	return th
}
