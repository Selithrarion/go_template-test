package main

import (
	"github.com/stretchr/testify/assert"
	"selithrarion/perfect-template/cmd/config"
	"testing"
)

func TestOfTest(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func LoadConfigTest(t *testing.T) {
	loadedConfig, err := config.LoadConfig("../..")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "postgres", loadedConfig.DBLogin)
}
