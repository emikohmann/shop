package config

import (
    "github.com/stretchr/testify/assert"
    "os"
    "strconv"
    "strings"
    "testing"
)

func TestRead(t *testing.T) {
    oldValues := make(map[string]string)
    testConfig := map[string]string{
        "HTTP_PORT": "8081",
    }

    // save old values and set test config
    for key, value := range testConfig {
        oldValues[key] = os.Getenv(key)
        assert.NoError(t, os.Setenv(key, value))
    }

    // preserve old values at end
    defer func() {
        for key, value := range oldValues {
            assert.NoError(t, os.Setenv(key, value))
        }
    }()

    // assert the test config
    config, err := Read()
    assert.NoError(t, err)
    assert.NotNil(t, config)
    assert.Equal(t, 8081, config.HTTP.Port)
}

func TestReadError(t *testing.T) {
    oldValues := make(map[string]string)
    testConfig := map[string]string{
        "HTTP_PORT": "invalid_value",
    }

    // save old values and set test config
    for key, value := range testConfig {
        oldValues[key] = os.Getenv(key)
        assert.NoError(t, os.Setenv(key, value))
    }

    // preserve old values at end
    defer func() {
        for key, value := range oldValues {
            assert.NoError(t, os.Setenv(key, value))
        }
    }()

    // assert the test config
    config, err := Read()
    assert.Nil(t, config)
    assert.Error(t, err)
    assert.True(t, strings.Contains(err.Error(), strconv.ErrSyntax.Error()))
}
