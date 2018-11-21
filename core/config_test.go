package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetConfig(t *testing.T) {
	config, err := SetConfig()
	assert.Nil(t, err)
	assert.NotNil(t,config.JWTKeyString)
}