package github_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("1234")
	assert.EqualValues(t, "token 1234", header)
}
