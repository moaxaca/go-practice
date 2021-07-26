package rest_server

import (
	"github.com/stretchr/testify/assert"
	"io.parcely.address_validation/test/util"
	"testing"
)

func TestAddressValidation(t *testing.T) {
	th := util.CreateTestRestHarness()
	resp, _ := th.Get("/validate")
	assert.Equal(t, resp.StatusCode, 200)
}
