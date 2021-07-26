package util

import (
	"io"
	"io.parcely.address_validation/internal"
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
	config := internal.RestServerConfiguration{}
	config.Name = "test suite"
	config.Address = ":8079"
	srv := internal.CreateRestServer(config)
	err := srv.Start()
	if err != nil {
		// PANIC
		return TestRestHarness{}
	}
	th := TestRestHarness{}
	th.Hostname = "localhost"+config.Name
	return th
}
