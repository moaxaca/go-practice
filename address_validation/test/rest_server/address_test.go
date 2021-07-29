package rest_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io.parcely.address_validation/internal/controllers"
	"io.parcely.address_validation/test/util"
	"testing"
)

func TestAddressValidation(t *testing.T) {
	th := util.CreateTestRestHarness()
	tests := []struct {
		status  int
		request controllers.AddressValidationRequest
	}{
		{status: 200, request: controllers.AddressValidationRequest {AddressLines: []string {"6272 Pacific Coast Hwy"}, Locality: "Long Beach", PostalCode: "90803", Region: "CA"}},
		//{status: "404", request: map[string]string {"addressLines": "test", "region": "test"}},
	}

	for _, tc := range tests {
		jsonValue, _ := json.Marshal(tc.request)
		requestJson := string(jsonValue)
		resp, _ := th.Post("/validate",  "application/json", bytes.NewBuffer([]byte(requestJson)))
		assert.Equal(t, resp.StatusCode, tc.status)
		b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b))
		resp.Body.Close()
	}
}
