package rest_test

import (
	"github.com/stretchr/testify/assert"
	"io.parcely.address_validation/test/util"
	"testing"
)

func TestExample(t *testing.T) {
	th := util.CreateTestRestHarness()
	resp, _ := th.Get("/example")
	assert.Equal(t, resp.StatusCode, 200)
}
